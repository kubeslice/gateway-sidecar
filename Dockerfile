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

ARG PLATFORM

FROM ${PLATFORM}/golang:1.17.7-alpine3.15 as gobuilder

# Install git.

# Git is required for fetching the dependencies.

RUN apk update && apk add --no-cache git make build-base

# Set the Go source path

WORKDIR /

COPY . .

# Build the binary.

RUN go mod download &&\
    go env -w GOPRIVATE=bitbucket.org/realtimeai && \
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o bin/avesha-kubeslice-gw-sidecar main.go


# Build reduced image from base alpine

FROM ${PLATFORM}/alpine:3.15

# tc - is needed for traffic control and shaping on the sidecar.  it is part of the iproute2

RUN apk add --no-cache ca-certificates \
    iproute2

# Copy our static executable.

COPY --from=gobuilder bin/avesha-kubeslice-gw-sidecar .

EXPOSE 5000

EXPOSE 8080

# Or could be CMD

ENTRYPOINT ["./avesha-kubeslice-gw-sidecar"]
