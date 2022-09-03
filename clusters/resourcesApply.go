package clusters

import (
	"context"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ApplyResources(clientset *kubernetes.Clientset, clientName string) {
	c, ok := GetClustersByName(clientName)
	if !ok {
		log.Printf("cluster %s is not registered, nothing to apply", clientName)
		return
	}

	log.Println("applying resources to cluster", clientName)
	ctx := context.TODO()
	cr := c.Resouces
	applyNamespaces(clientset, &cr, ctx)
	applyServiceAccounts(clientset, &cr, ctx)
	applyRoles(clientset, &cr, ctx)
	applyRoleBindings(clientset, &cr, ctx)
	applyPods(clientset, &cr, ctx)
	applyDeployments(clientset, &cr, ctx)
	applyServices(clientset, &cr, ctx)
}

func applyNamespaces(clientset *kubernetes.Clientset, cr *SupportResources, ctx context.Context) {
	for _, ns := range cr.V1Namespace {
		_, err := clientset.CoreV1().Namespaces().Get(ctx, ns.Name, metav1.GetOptions{})
		logResourceGet(ns, err)
		if err != nil {
			_, err := clientset.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
			logResourceCreate(ns, err)
		} else {
			_, err := clientset.CoreV1().Namespaces().Update(ctx, ns, metav1.UpdateOptions{})
			logResourceUpdate(ns, err)
		}
	}
}

func applyPods(clientset *kubernetes.Clientset, cr *SupportResources, ctx context.Context) {
	for _, p := range cr.V1Pods {
		_, err := clientset.CoreV1().Pods(namespaceAdapter(p.Namespace)).Get(ctx, p.Name, metav1.GetOptions{})
		logResourceGet(p, err)
		if err != nil {
			_, err := clientset.CoreV1().Pods(namespaceAdapter(p.Namespace)).Create(ctx, p, metav1.CreateOptions{})
			logResourceCreate(p, err)
		} else {
			_, err := clientset.CoreV1().Pods(namespaceAdapter(p.Namespace)).Update(ctx, p, metav1.UpdateOptions{})
			logResourceUpdate(p, err)
		}
	}
}

func applyDeployments(clientset *kubernetes.Clientset, cr *SupportResources, ctx context.Context) {
	for _, d := range cr.AppV1Deployments {
		_, err := clientset.AppsV1().Deployments(namespaceAdapter(d.Namespace)).Get(ctx, d.Name, metav1.GetOptions{})
		logResourceGet(d, err)
		if err != nil {
			_, err := clientset.AppsV1().Deployments(namespaceAdapter(d.Namespace)).Create(ctx, d, metav1.CreateOptions{})
			logResourceCreate(d, err)
		} else {
			_, err := clientset.AppsV1().Deployments(namespaceAdapter(d.Namespace)).Update(ctx, d, metav1.UpdateOptions{})
			logResourceUpdate(d, err)
		}
	}
}

func applyServices(clientset *kubernetes.Clientset, cr *SupportResources, ctx context.Context) {
	for _, s := range cr.V1Services {
		_, err := clientset.CoreV1().Services(namespaceAdapter(s.Namespace)).Get(ctx, s.Name, metav1.GetOptions{})
		logResourceGet(s, err)
		if err != nil {
			_, err := clientset.CoreV1().Services(namespaceAdapter(s.Namespace)).Create(ctx, s, metav1.CreateOptions{})
			logResourceCreate(s, err)
		} else {
			_, err := clientset.CoreV1().Services(namespaceAdapter(s.Namespace)).Update(ctx, s, metav1.UpdateOptions{})
			logResourceUpdate(s, err)
		}
	}
}

func applyServiceAccounts(clientset *kubernetes.Clientset, cr *SupportResources, ctx context.Context) {
	for _, sa := range cr.V1ServiceAccount {
		_, err := clientset.CoreV1().ServiceAccounts(namespaceAdapter(sa.Namespace)).Get(ctx, sa.Name, metav1.GetOptions{})
		logResourceGet(sa, err)
		if err != nil {
			_, err := clientset.CoreV1().ServiceAccounts(namespaceAdapter(sa.Namespace)).Create(ctx, sa, metav1.CreateOptions{})
			logResourceCreate(sa, err)
		} else {
			_, err := clientset.CoreV1().ServiceAccounts(namespaceAdapter(sa.Namespace)).Update(ctx, sa, metav1.UpdateOptions{})
			logResourceUpdate(sa, err)
		}
	}
}

func applyRoles(clientset *kubernetes.Clientset, cr *SupportResources, ctx context.Context) {
	for _, r := range cr.RbacV1Role {
		_, err := clientset.RbacV1().Roles(namespaceAdapter(r.Namespace)).Get(ctx, r.Name, metav1.GetOptions{})
		logResourceGet(r, err)
		if err != nil {
			_, err := clientset.RbacV1().Roles(namespaceAdapter(r.Namespace)).Create(ctx, r, metav1.CreateOptions{})
			logResourceCreate(r, err)
		} else {
			_, err := clientset.RbacV1().Roles(namespaceAdapter(r.Namespace)).Update(ctx, r, metav1.UpdateOptions{})
			logResourceUpdate(r, err)
		}
	}
}

func applyRoleBindings(clientset *kubernetes.Clientset, cr *SupportResources, ctx context.Context) {
	for _, rb := range cr.RbacV1RoleBinding {
		_, err := clientset.RbacV1().RoleBindings(namespaceAdapter(rb.Namespace)).Get(ctx, rb.Name, metav1.GetOptions{})
		logResourceGet(rb, err)
		if err != nil {
			_, err := clientset.RbacV1().RoleBindings(namespaceAdapter(rb.Namespace)).Create(ctx, rb, metav1.CreateOptions{})
			logResourceCreate(rb, err)
		} else {
			_, err := clientset.RbacV1().RoleBindings(namespaceAdapter(rb.Namespace)).Update(ctx, rb, metav1.UpdateOptions{})
			logResourceUpdate(rb, err)
		}
	}
}

func logResourceGet(r metav1.Object, err error) {
	if err != nil {
		log.Printf("failed to get %T %v %v", r, r.GetName(), err)
	} else {
		log.Printf("succeed to get %T %v", r, r.GetName())
	}
}

func logResourceCreate(r metav1.Object, err error) {
	if err != nil {
		log.Printf("failed to create %T %v %v", r, r.GetName(), err)
	} else {
		log.Printf("%T %v created", r, r.GetName())
	}
}

func logResourceUpdate(r metav1.Object, err error) {
	if err != nil {
		log.Printf("failed to update %T %v %v", r, r.GetName(), err)
	} else {
		log.Printf("%T %v updated", r, r.GetName())
	}
}

func namespaceAdapter(ns string) string {
	if ns == "" {
		return v1.NamespaceDefault
	}
	return ns
}
