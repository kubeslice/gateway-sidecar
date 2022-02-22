ARG PLATFORM

FROM ${PLATFORM}/golang:1.17.7-alpine3.15 as gobuilder

LABEL maintainer="Avesha Systems LLC"

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
