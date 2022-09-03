package hello_busybox

import (
	"github.com/cugtyt/grimoire/clusters"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	serviceName = "hello-busybox"
	targetClusters = "Minikube"
	serviceAccount = "hello-serviceaccount"
	namespace = "hello-busybox"
	role = "hello-role"
	roleBinding = "hello-rolebinding"
)

func init() {
	targetClusters := clusters.GetClustersByLabel([]string{"minikube"})
	clusters.AddResources(targetClusters, []metav1.Object{
		helloServices(), 
		helloDeployment(),
		helloNamespace(),
	})
}
