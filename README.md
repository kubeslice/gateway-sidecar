# gateway-sidecar

gateway-sidecar usages kubernetes sidecar pattern, for establishing connection between 2 clusters using OpenVPN pods.

## Getting Started

It is strongly recommended to use a released version.

### Prerequisites

* Docker installed and running in your local machine
* A running [`kind`](https://kind.sigs.k8s.io/) or [`Docker Desktop Kubernetes`](https://docs.docker.com/desktop/kubernetes/)
  cluster 
* [`kubectl`](https://kubernetes.io/docs/tasks/tools/) installed and configured

### Installation
To install: 

1. Clone the latest version of gateway-sidecar from  the `master` branch.

```bash
  git clone https://github.com/kubeslice/gateway-sidecar.git
  cd gateway-sidecar
```

## License

This project is released under the Apache 2.0 License.
