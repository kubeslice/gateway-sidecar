# VERSION defines the project version for the bundle.
# Update this value when you upgrade the version of your project.
VERSION ?= latest-stable

IMG ?= docker.io/aveshasystems/kubeslice-gw-sidecar:$(VERSION)

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
docker-build: ## Build docker image with the manager.
	docker buildx create --name container --driver=docker-container || true
	docker build --builder container --platform linux/amd64,linux/arm64 -t ${IMG} .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	docker buildx create --name container --driver=docker-container || true
	docker build --push --builder container --platform linux/amd64,linux/arm64 -t ${IMG} .

.PHONY: chart-deploy
chart-deploy:
	## Deploy the artifacts using helm
	## Usage: make chart-deploy VALUESFILE=[valuesfilename]
	helm upgrade --install kubeslice-worker -n kubeslice-system avesha/kubeslice-worker -f ${VALUESFILE}
