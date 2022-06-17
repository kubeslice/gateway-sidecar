# Development guidelines for Router Sidecar

The Slice VPN Gateway is a slice network service component that provides a secure VPN tunnel between any two clusters that are a part of the slice.

## Building and Installing `gateway-sidecar` in a Local Kind Cluster
For more information, see [getting started with kind clusters](https://docs.avesha.io/opensource/getting-started-with-kind-clusters).

### Setting up Development Environment

* Go (version 1.17 or later) installed and configured in your machine ([Installing Go](https://go.dev/dl/))
* Docker installed and running in your local machine
* A running [`kind`](https://kind.sigs.k8s.io/)  cluster
* [`kubectl`](https://kubernetes.io/docs/tasks/tools/) installed and configured
* Follow the getting started from above, to install [kubeslice-controller](https://github.com/kubeslice/kubeslice-controller) 



### Building Docker Images

1. Clone the latest version of kubeslice-controller from  the `master` branch.

```bash
git clone https://github.com/kubeslice/gateway-sidecar.git
cd gateway-sidecar
```

2. Adjust image name variable `IMG` in the [`Makefile`](Makefile) to change the docker tag to be built.
   Default image is set as `IMG ?= aveshasystems/kubeslice-gw-sidecar:${VERSION}`. Modify this if required.

```bash
make docker-build
```
### Running Local Image on Kind Clusters

1. You can load the gateway-sidecar docker image into the kind cluster.

```bash
kind load docker-image my-custom-image:unique-tag --name clustername
```

Example

```console
kind load docker-image aveshasystems/kubeslice-gw-sidecar:1.2.1 --name kind
```

2. Check the loaded image in the cluster. Modify the node name if required.

```console
docker exec -it <node-name> crictl images
```

Example.

```console
docker exec -it kind-control-plane crictl images
```

### Deploy in a Cluster

Update chart values file `yourvaluesfile.yaml` that you have previously created.
Refer to [values.yaml](https://github.com/kubeslice/charts/blob/master/charts/kubeslice-worker/values.yaml) to create `yourvaluesfiel.yaml` and update the routerSidecar image subsection to use the local image.

From the sample:

```
routerSidecar:
  image: docker.io/aveshasystems/gw-sidecar
  tag: 0.1.0
```

Change it to:

```
routerSidecar:
  image: <my-custom-image>
  tag: <unique-tag>
```

Deploy the Updated Chart

```console
make chart-deploy VALUESFILE=yourvaluesfile.yaml
```

### Verify the gateway-sidecar Pods are Running

```bash
kubectl describe pod <gateway pod name> -n kubeslice-system
```

## License

Apache License 2.0
