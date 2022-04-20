.PHONY: proto
proto: ##generate the proto files
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/sidecar/sidecarpb/gw_sidecar.proto

.PHONY: fmt
fmt: ##run go fmt
	go fmt ./...

.PHONY: build
build: fmt
	go build -o bin/gw-sidecar main.go

.PHONY: docker-build
docker-build: build
	docker build -t kubeslice-gw-sidecar:latest-release --build-arg PLATFORM=amd64 . && docker tag kubeslice-gw-sidecar:latest-release nexus.dev.aveshalabs.io/kubeslice-gw-sidecar:latest-stable

.PHONY: docker-push
docker-push:
	docker push nexus.dev.aveshalabs.io/kubeslice-gw-sidecar:latest-stable