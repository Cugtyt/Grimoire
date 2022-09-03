package main

import (
	"os"

	"github.com/cugtyt/grimoire/clusters"
	_ "github.com/cugtyt/grimoire/services"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	clusters.ApplyResources(clientset, os.Getenv("CLUSTER_NAME"))
}
