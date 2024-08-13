package listener

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"modi/client"
	"modi/src/service"
)

func InitInformerListener() {
	fact := informers.NewSharedInformerFactory(client.K8sClient, 0)

	nsInformer := fact.Core().V1().Namespaces().Informer()
	_, err := nsInformer.AddEventHandler(&service.NamespaceHandler{})
	if err != nil {
		return
	}

	// 初始化 deployment监听
	depInformer := fact.Apps().V1().Deployments().Informer()
	_, err = depInformer.AddEventHandler(&service.DeploymentHandler{})
	if err != nil {
		return
	}

	// 初始化 replicaSet监听
	rsInformer := fact.Apps().V1().ReplicaSets().Informer()
	_, err = rsInformer.AddEventHandler(&service.RsHandler{})
	if err != nil {
		return
	}

	// 初始化 pod监听
	podInformer := fact.Core().V1().Pods().Informer()
	_, err = podInformer.AddEventHandler(&service.PodHandler{})
	if err != nil {
		return
	}

	// 初始化 event 监听
	eventInformer := fact.Core().V1().Events().Informer()
	_, err = eventInformer.AddEventHandler(&service.EventHandler{})
	if err != nil {
		return
	}

	// 初始化 secret 监听
	secretInformer := fact.Core().V1().Secrets().Informer()
	_, err = secretInformer.AddEventHandler(&service.SecretHandler{})
	if err != nil {
		return
	}

	// 初始化configMap 监听
	configMapInformer := fact.Core().V1().ConfigMaps().Informer()
	_, err = configMapInformer.AddEventHandler(&service.ConfigMapHandler{})
	if err != nil {
		return
	}

	fact.Start(wait.NeverStop)
}
