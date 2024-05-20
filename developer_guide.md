# Development Guidelines for Router Sidecar

The Slice VPN Gateway is a slice network service component that provides a secure VPN tunnel between any two clusters that are a part of the slice.

## Building and Installing gateway-sidecar in a Local Kind Cluster
For more information, see [getting started with clusters](https://kubeslice.io/documentation/open-source/1.2.0/category/get-started).

### Setting up Development Environment
Ensure that: 
* Go (version 1.17 or later) is installed and configured in your machine ([Installing Go](https://go.dev/dl/)).
* Docker is installed and running in your local machine
* You have a running [kind](https://kind.sigs.k8s.io/) cluster
* [kubectl](https://kubernetes.io/docs/tasks/tools/)  is installed and configured on the cluster.
* You follow the getting started from above to install [kubeslice-controller](https://github.com/kubeslice/kubeslice-controller).

### Building Docker Images

1. Clone the latest version of kubeslice-controller from  the master branch.

```
git clone https://github.com/kubeslice/gateway-sidecar.git
cd gateway-sidecar
```


2. Adjust image name variable IMG in the [Makefile](Makefile) to change the docker tag to be built.
   The default image is set as IMG ?= aveshasystems/kubeslice-gw-sidecar:${VERSION}. Modify this if required.
```
make docker-build
```
### Running Local Image on Kind Clusters

1. You can load the gateway-sidecar docker image into the kind cluster.
```
kind load docker-image my-custom-image:unique-tag --name clustername
```

Example
```
kind load docker-image aveshasystems/kubeslice-gw-sidecar:1.2.1 --name kind
```

2. Check the loaded image in the cluster. Modify the node name if required.
```
docker exec -it <node-name> crictl images
```

Example.
```
docker exec -it kind-control-plane crictl images
```

### Deploy in a Cluster

Update chart values file yourvaluesfile.yaml that you have previously created.
Refer to [values.yaml](https://github.com/kubeslice/charts/blob/master/charts/kubeslice-worker/values.yaml) to create yourvaluesfile.yaml and update the routerSidecar image subsection to use the local image.

From the sample:
Change the following parameter values
```
gatewaySidecar:
  image: docker.io/aveshasystems/gw-sidecar
  tag: 0.1.0
```

Change them to:

```
gatewaySidecar:
  image: <my-custom-image>
  tag: <unique-tag>
```

Deploy the Updated Chart
```
make chart-deploy VALUESFILE=yourvaluesfile.yaml
```

### Verify the gateway-sidecar Pods are Running
```
kubectl describe pod <gateway pod name> -n kubeslice-system
```
### Uninstalling the kubeslice-worker

Refer to the [uninstallation guide](https://kubeslice.io/documentation/open-source/1.2.0/uninstall-kubeslice/).

1. [Offboard](https://kubeslice.io/documentation/open-source/1.2.0/uninstall-kubeslice/#offboard-application-namespaces) the namespaces from the slice.

2. [Delete](https://kubeslice.io/documentation/open-source/1.2.0/uninstall-kubeslice/#delete-slices) the slice.

3. On the worker cluster, undeploy the kubeslice-worker charts.


# uninstall all the resources
```
make chart-undeploy
```

## License

Apache License 2.0
