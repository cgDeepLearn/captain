## Proposals

To make convenient for invocating of CRD client operation methods,we recommend the CRD that Captain integrated implement a uniform CRD client interface.

Interface invocation consult **Usage**.

Interface implementation consult **Best Practice**.

## Usage

```
import (
	"context"
	"k8s.io/klog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"captain/pkg/simple/client/k8s"
)
...
func demo(){
	...
	// init client
	kubernetesClient, err := k8s.NewKubernetesClient(s.KubernetesOptions)
	if err != nil {
		klog.Errorf("Failed to create kubernetes clientset %v", err)
		return err
	}
	...
	// Interface invocation: get gaiaCluster
	kubernetesClient.Crd().V1alpha1().GaiaCluster(namespace).Get(context.TODO(), clusterName, metav1.GetOptions{})
	// Interface implementation: get cluster
	kubernetesClient.Crd().V1alpha1().Clusters().Get(context.TODO(), clusterName, metav1.GetOptions{})
	...
}
```

## Best Practice

Package directory

```
captain
|-pkg
  |-crd
    |-v1beta1
      |-v1beta1Client.go
    |-clientSet.go
```

1. ADD your crd's getter interface in `V1beta1Interface `Interface，for example adding `ClustersGetter` which is generated by [code-generate](https://github.com/kubernetes/code-generator)(refer:`captain/pkg/client/clientset/versioned/typed/cluster/v1alpha1.go ClustersGetter`)

```
type V1beta1Interface interface {
	gaia.GaiaClusterGetter
	gaia.GaiaNodeGetter
	gaia.GaiaSetGetter
	// add cluster getter
	clusterv1alpha1.ClustersGetter
	// todo list, add your getter
}
```

2. IMPLATEMENT your getter function with `V1beta1Client`, Versioned is a fixed referencefixed refernce that generated by [code-generate](https://github.com/kubernetes/code-generator)(refer:`captain/pkg/client/clientset/versioned/clientset.go Clientset`). eg.

```
   func (c *V1beta1Client) Clusters() clusterv1alpha1.ClusterInterface {
	return c.Versioned.ClusterV1beta1().Clusters()
   }
```

Now, you can easily use chain call.

```
kubernetesClient.Crd().V1beta1().Clusters().Get(context.TODO(), clusterName, metav1.GetOptions{})
```