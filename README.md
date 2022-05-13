# gateway-sidecar

![Docker Image Size](https://img.shields.io/docker/image-size/aveshasystems/gw-sidecar/latest)
![Docker Image Version](https://img.shields.io/docker/v/aveshasystems/gw-sidecar)

The Slice VPN Gateway is a slice network service component that provides a secure VPN tunnel between any two clusters that are a part of the slice. 

## Getting Started

[TBD: Add link to getting started] 
It is strongly recommended to use a released version.

### Prerequisites

* Docker installed and running in your local machine
* A running [`kind`](https://kind.sigs.k8s.io/)
* [`kubectl`](https://kubernetes.io/docs/tasks/tools/) installed and configured
* Follow the getting started from above, to install [`kubeslice-controller`](https://github.com/kubeslice/kubeslice-controller) and [`worker-operator`](https://github.com/kubeslice/worker-operator)

# Local build and update 

## Latest docker image
[TBD link to docker hub]

## Setting up your helm repo

If you have not added avesha helm repo yet, add it

```console
helm repo add avesha https://kubeslice.github.io/charts/
```

upgrade the avesha helm repo

```console
helm repo update

### Build docker images

```bash
git clone https://github.com/kubeslice/gateway-sidecar.git
cd gateway-sidecar
make docker-build
```

### Running locally on Kind
You can load the gateway-sidecar docker image into kind cluster

```bash
kind load docker-image my-custom-image:unique-tag --name clustername
```

### Verification
You can view the sidecar container by describing the gateway pod: 

```bash
kubectl describe pod <gateway pod name> -n kubeslice-system
```

## License
Apache 2.0 License.
