package client

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"modi/internal/kind"
)

func InitInformerListener() {
	fact := informers.NewSharedInformerFactory(K8sClient, 0)

	nsInformer := fact.Core().V1().Namespaces().Informer()
	_, err := nsInformer.AddEventHandler(&kind.NamespaceHandler{})
	if err != nil {
		return
	}

	// 初始化 deployment监听
	depInformer := fact.Apps().V1().Deployments().Informer()
	_, err = depInformer.AddEventHandler(&kind.DeploymentHandler{})
	if err != nil {
		return
	}

	// 初始化 replicaSet监听
	rsInformer := fact.Apps().V1().ReplicaSets().Informer()
	_, err = rsInformer.AddEventHandler(&kind.RsHandler{})
	if err != nil {
		return
	}

	// 初始化 pod监听
	podInformer := fact.Core().V1().Pods().Informer()
	_, err = podInformer.AddEventHandler(&kind.PodHandler{})
	if err != nil {
		return
	}

	// 初始化 event 监听
	eventInformer := fact.Core().V1().Events().Informer()
	_, err = eventInformer.AddEventHandler(&kind.EventHandler{})
	if err != nil {
		return
	}

	// 初始化 secret 监听
	secretInformer := fact.Core().V1().Secrets().Informer()
	_, err = secretInformer.AddEventHandler(&kind.SecretHandler{})
	if err != nil {
		return
	}

	fact.Start(wait.NeverStop)
}
