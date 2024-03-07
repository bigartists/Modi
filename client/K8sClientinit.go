package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

var K8sClient *kubernetes.Clientset

func init() {
	kubeconfig := "etc/ai-stage.yaml"
	//kubeconfig := "etc/ai-dx-test.yaml"

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	K8sClient = clientset
}
