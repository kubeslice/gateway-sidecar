# Gateway Sidecar

The Slice VPN Gateway is a slice network service component that provides a secure VPN tunnel between any two clusters that are a part of the slice. 

## Get Started

Please refer to our documentation on:
- [Install KubeSlice on cloud clusters](https://kubeslice.io/documentation/open-source/0.6.0/getting-started-with-cloud-clusters/installing-kubeslice/installing-the-kubeslice-controller)
- [Install KubeSlice on kind clusters](https://kubeslice.io/documentation/open-source/0.6.0/tutorials/kind-install-kubeslice-controller)

### Prerequisites
Before you begin, make sure the following prerequisites are met:
* Docker is installed and running on your local machine.
* A running [`kind`](https://kind.sigs.k8s.io/).
* [`kubectl`](https://kubernetes.io/docs/tasks/tools/) is installed and configured.
* You have prepared the environment to install [`kubeslice-controller`](https://github.com/kubeslice/kubeslice-controller) on the controller cluster
 and [`worker-operator`](https://github.com/kubeslice/worker-operator) on the worker cluster. For more information, see [Prerequisites](https://kubeslice.io/documentation/open-source/0.6.0/getting-started-with-cloud-clusters/prerequisites/).

# Build and Deploy Gateway Sidecar on a Kind Cluster 

To download the latest gateway-sidecar docker hub image, click [here](https://hub.docker.com/r/aveshasystems/gw-sidecar).

The following command pulls the latest docker image:

```console
docker pull aveshasystems/gw-sidecar:latest
```

## Set up Your Helm Repo

If you have not added avesha helm repo yet, add it.

```console
helm repo add avesha https://kubeslice.github.io/charts/
```

Upgrade the avesha helm repo.

```console
helm repo update
```
### Build Docker Images

1. Clone the latest version of gateway sidecar from  the `master` branch.

   ```bash
   git clone https://github.com/kubeslice/gateway-sidecar.git
   cd gateway-sidecar
   ```

2. Edit the `VERSION` variable in the Makefile to change the docker tag to be built.
   The image is set as `docker.io/aveshasystems/kubeslice-gw-sidecar:$(VERSION)` in the Makefile. Change this if required

   ```
   make docker-build
   ```

### Run Locally on a Kind Cluster
1. You can load the gateway-sidecar docker image into a kind cluster.

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

   
   Example

   ```console
   docker exec -it kind-control-plane crictl images
   ```

### Deploy Gateway Sidecar on a Cluster

Update the chart values file called `yourvaluesfile.yaml` that you have previously created.
Refer to the [values.yaml](https://github.com/kubeslice/charts/blob/master/charts/kubeslice-worker/values.yaml) to create `yourvaluesfiel.yaml` and update the gateway-sidecar image subsection to use the local image.

From the sample:

```
gateway:
  image: docker.io/aveshasystems/gw-sidecar
  tag: 0.1.0
```

Change it to:

```
gateway:
  image: <my-custom-image>
  tag: <unique-tag>
```

Deploy the Updated Chart

```console
make chart-deploy VALUESFILE=yourvaluesfile.yaml
```

### Verify the Installation

Verify the Gateway Sidecar Container is running by checking the status of gateway pod belonging to the `kubeslice-system` namespace.

```bash
kubectl describe pod <gateway pod name> -n kubeslice-system
```

## License
Apache 2.0 License.
