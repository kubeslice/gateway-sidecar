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

### Usages
You can view the sidecar container by describing the gateway pod: 

```bash
kubectl describe pod <gateway pod name> -n kubeslice-system
```

## License
This project is released under the Apache 2.0 License.
