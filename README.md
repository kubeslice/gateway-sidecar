# gateway-sidecar

The Slice VPN Gateway is a slice network service component that provides a secure VPN tunnel between any two clusters that are a part of the slice. 

## Getting Started

It is strongly recommended to use a released version.

### Prerequisites

* Docker installed and running in your local machine
* A running [`kind`](https://kind.sigs.k8s.io/) or [`Docker Desktop Kubernetes`](https://docs.docker.com/desktop/kubernetes/)
  cluster 
* [`kubectl`](https://kubernetes.io/docs/tasks/tools/) installed and configured
* Install the [`operator`](https://github.com/kubeslice/operator)

### Build docker images

```bash
git clone https://github.com/kubeslice/gateway-sidecar.git
cd gateway-sidecar
make docker-build
```

### Running locally on Kind
Load the docker image into kind cluster

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
