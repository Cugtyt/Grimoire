package hello_busybox

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appv1 "k8s.io/api/apps/v1"
)

func helloDeployment() *appv1.Deployment {
	var replicas int32 = 1
	return &appv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
			Namespace: namespace,
			Labels: map[string]string{
				"app": serviceName,
			},
		},
		Spec: appv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": serviceName,
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": serviceName,
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  serviceName,
							Image: "busybox",
							Args:  []string{"echo", "hello world"},
						},
					},
				},
			},
		},
	}
}