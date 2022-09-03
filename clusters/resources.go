package clusters

import (
	"log"

	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SupportResources struct {
	V1Services        []*v1.Service
	AppV1Deployments  []*appv1.Deployment
	V1Pods            []*v1.Pod
	V1Namespace       []*v1.Namespace
	V1ServiceAccount  []*v1.ServiceAccount
	RbacV1Role        []*rbacv1.Role
	RbacV1RoleBinding []*rbacv1.RoleBinding
}

func InitResources() SupportResources {
	return SupportResources{
		V1Services:        make([]*v1.Service, 0),
		AppV1Deployments:  make([]*appv1.Deployment, 0),
		V1Pods:            make([]*v1.Pod, 0),
		V1Namespace:       make([]*v1.Namespace, 0),
		V1ServiceAccount:  make([]*v1.ServiceAccount, 0),
		RbacV1Role:        make([]*rbacv1.Role, 0),
		RbacV1RoleBinding: make([]*rbacv1.RoleBinding, 0),
	}
}

func AddResources(targetClusters []*Cluster, resources []metav1.Object) {
	for i, c := range targetClusters {
		cr := c.Resouces
		for _, r := range resources {
			log.Printf("adding %T %v to cluster %s\n", r, r.GetName(), targetClusters[i].Name)
			switch r := r.(type) {
			case *v1.Service:
				targetClusters[i].Resouces.V1Services = append(cr.V1Services, r)
			case *appv1.Deployment:
				targetClusters[i].Resouces.AppV1Deployments = append(cr.AppV1Deployments, r)
			case *v1.Pod:
				targetClusters[i].Resouces.V1Pods = append(cr.V1Pods, r)
			case *v1.Namespace:
				targetClusters[i].Resouces.V1Namespace = append(cr.V1Namespace, r)
			case *v1.ServiceAccount:
				targetClusters[i].Resouces.V1ServiceAccount = append(cr.V1ServiceAccount, r)
			case *rbacv1.Role:
				targetClusters[i].Resouces.RbacV1Role = append(cr.RbacV1Role, r)
			case *rbacv1.RoleBinding:
				targetClusters[i].Resouces.RbacV1RoleBinding = append(cr.RbacV1RoleBinding, r)
			}
		}
	}
}
