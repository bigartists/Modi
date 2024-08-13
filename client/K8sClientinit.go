package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

var K8sClient *kubernetes.Clientset
var K8sClientRestConfig *rest.Config

func init() {
	//kubeconfig := "etc/ai-stage.yaml"
	kubeconfig := "etc/ai-dx-test.yaml"
	//kubeconfig := "etc/npu-910b-test.yaml"

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	K8sClient = clientset
	K8sClientRestConfig = config
}
