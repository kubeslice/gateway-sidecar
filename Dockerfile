# syntax=docker/dockerfile:1.4
##########################################################
#Dockerfile
#Copyright (c) 2022 Avesha, Inc. All rights reserved.
#
#SPDX-License-Identifier: Apache-2.0
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.
##########################################################

# Build stage - use BUILDPLATFORM for native compilation
FROM --platform=$BUILDPLATFORM golang:1.24-alpine3.21 AS gobuilder

# Multi-arch build arguments
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

# Install build dependencies
# Git is required for fetching the dependencies.
# build-base is needed for CGO compilation
# For cross-compilation, we need the target architecture's toolchain
RUN apk update && apk add --no-cache git make build-base && \
    if [ "$TARGETARCH" != "amd64" ]; then \
        apk add --no-cache gcc-aarch64-linux-gnu g++-aarch64-linux-gnu || \
        apk add --no-cache crossbuild-essential-arm64 || \
        echo "Warning: ARM64 cross-compilation toolchain not available via apk"; \
    fi

WORKDIR /workspace

# Copy go mod files first for better layer caching
COPY go.mod go.sum ./

# Download dependencies with cache mount for better performance
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && go mod verify

# Copy source code
COPY . .

# Build the binary with cross-compilation support
# CGO_ENABLED=1 is required for some dependencies (e.g., netlink, syscall)
# For ARM64 cross-compilation, we need to set CC and CXX to the cross-compiler
RUN --mount=type=cache,target=/root/.cache/go-build \
    if [ "$TARGETARCH" = "arm64" ]; then \
        export CC=aarch64-linux-gnu-gcc && \
        export CXX=aarch64-linux-gnu-g++ && \
        export CGO_ENABLED=1 && \
        export GOOS=${TARGETOS} && \
        export GOARCH=${TARGETARCH} && \
        go build -a -trimpath -ldflags="-w -s" -o bin/kubeslice-gw-sidecar main.go; \
    else \
        CGO_ENABLED=1 \
        GOOS=${TARGETOS} \
        GOARCH=${TARGETARCH} \
        go build -a -trimpath -ldflags="-w -s" -o bin/kubeslice-gw-sidecar main.go; \
    fi

# Final stage - use TARGETPLATFORM for correct base image
FROM --platform=$TARGETPLATFORM alpine:3.21

# Multi-arch build arguments
ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH

LABEL maintainer="Avesha Systems" \
      org.opencontainers.image.title="kubeslice-gw-sidecar" \
      org.opencontainers.image.description="KubeSlice Gateway Sidecar for multi-architecture support" \
      org.opencontainers.image.architecture=${TARGETARCH} \
      org.opencontainers.image.os=${TARGETOS}

# tc - is needed for traffic control and shaping on the sidecar. it is part of the iproute2
RUN apk add --no-cache ca-certificates iproute2 && \
    rm -rf /var/cache/apk/*

# Copy our static executable from builder stage
COPY --from=gobuilder /workspace/bin/kubeslice-gw-sidecar /kubeslice-gw-sidecar

# Set executable permissions
RUN chmod +x /kubeslice-gw-sidecar

EXPOSE 5000 8080

# Use exec form for better signal handling
ENTRYPOINT ["/kubeslice-gw-sidecar"]
