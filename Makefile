# VERSION defines the project version for the bundle.
# Update this value when you upgrade the version of your project.
VERSION ?= latest-stable

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
	docker build -t kubeslice-gw-sidecar:${VERSION} --build-arg PLATFORM=amd64 . && docker tag kubeslice-gw-sidecar:${VERSION} docker.io/aveshasystems/kubeslice-gw-sidecar:${VERSION}

.PHONY: docker-push
docker-push:
	docker push docker.io/aveshasystems/kubeslice-gw-sidecar:${VERSION}

.PHONY: chart-deploy
chart-deploy:
	## Deploy the artifacts using helm
	## Usage: make chart-deploy VALUESFILE=[valuesfilename]
	helm upgrade --install kubeslice-worker -n kubeslice-system avesha/kubeslice-worker -f ${VALUESFILE}
