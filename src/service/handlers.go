package service

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"log"
)

type DeploymentHandler struct {
}

func (d DeploymentHandler) OnAdd(obj interface{}, isInInitialList bool) {
	DeploymentMapInstance.Add(obj.(*v1.Deployment))
}

func (d DeploymentHandler) OnUpdate(oldObj, newObj interface{}) {
	err := DeploymentMapInstance.Update(newObj.(*v1.Deployment))
	if err != nil {
		return
	}
}

func (d DeploymentHandler) OnDelete(obj interface{}) {
	DeploymentMapInstance.Delete(obj.(*v1.Deployment))
}

type EventHandler struct {
}

func (this *EventHandler) storeData(obj interface{}, isDelete bool) {
	if event, ok := obj.(*corev1.Event); ok {
		key := fmt.Sprintf("%s_%s_%s", event.Namespace, event.InvolvedObject.Kind, event.InvolvedObject.Name)
		//if event.Namespace == "groot" {
		//	fmt.Println(".....event.Message.......", "key=", key, event.Message)
		//}
		//if key == "groot_Pod_ng1-5cb776c5f-wq4ld" {
		//	fmt.Println(".....event.Message.......", "key=", key, "eventType=", event.Type, "value=", event.Message)
		//}

		if !isDelete {
			EventMapInstance.Store(key, event)
		} else {
			EventMapInstance.Delete(key)
		}
	}
}

func (this EventHandler) OnAdd(obj interface{}, isInInitialList bool) {
	this.storeData(obj, false)
}

func (this EventHandler) OnUpdate(oldObj, newObj interface{}) {
	this.storeData(newObj, false)
}

func (this EventHandler) OnDelete(obj interface{}) {
	this.storeData(obj, true)
}

type NamespaceHandler struct {
}

func (nh NamespaceHandler) OnAdd(obj interface{}, isInInitialList bool) {
	NamespaceMapInstance.Add(obj.(*corev1.Namespace))
}

func (nh NamespaceHandler) OnUpdate(oldObj, newObj interface{}) {
	err := NamespaceMapInstance.Update(newObj.(*corev1.Namespace))
	if err != nil {
		return
	}
}

func (nh NamespaceHandler) OnDelete(obj interface{}) {
	namespace := obj.(*corev1.Namespace)
	NamespaceMapInstance.Delete(namespace)
}

type RsHandler struct {
}

func (r RsHandler) OnAdd(obj interface{}, isInInitialList bool) {
	RsMapInstance.Add(obj.(*v1.ReplicaSet))
}

func (r RsHandler) OnUpdate(oldObj, newObj interface{}) {
	err := RsMapInstance.Update(newObj.(*v1.ReplicaSet))
	if err != nil {
		log.Println(err)
	}
}

func (r RsHandler) OnDelete(obj interface{}) {
	if d, ok := obj.(*v1.ReplicaSet); ok {
		RsMapInstance.Delete(d)
	}
}

type PodHandler struct{}

func (p PodHandler) OnAdd(obj interface{}, isInInitialList bool) {
	PodMapInstance.Add(obj.(*corev1.Pod))
}

func (p PodHandler) OnUpdate(oldObj, newObj interface{}) {
	err := PodMapInstance.Update(newObj.(*corev1.Pod))
	if err != nil {
		log.Println(err)
	}
}

func (p PodHandler) OnDelete(obj interface{}) {
	if d, ok := obj.(*corev1.Pod); ok {
		PodMapInstance.Delete(d)
	}
}

type SecretHandler struct {
}

func (this SecretHandler) OnAdd(obj interface{}, isInInitialList bool) {
	SecretMapInstance.Add(obj.(*corev1.Secret))
}

func (this SecretHandler) OnUpdate(oldObj, newObj interface{}) {
	err := SecretMapInstance.Update(newObj.(*corev1.Secret))
	if err != nil {
		return
	}
}

func (this SecretHandler) OnDelete(obj interface{}) {
	SecretMapInstance.Delete(obj.(*corev1.Secret))
}

type ConfigMapHandler struct {
}

func (this *ConfigMapHandler) OnAdd(obj interface{}, isInInitialList bool) {
	ConfigMapInstance.Add(obj.(*corev1.ConfigMap))
}
func (this *ConfigMapHandler) OnUpdate(oldObj, newObj interface{}) {
	err := ConfigMapInstance.Update(newObj.(*corev1.ConfigMap))
	if err == false {
		return
	}
}
func (this *ConfigMapHandler) OnDelete(obj interface{}) {
	ConfigMapInstance.Delete(obj.(*corev1.ConfigMap))
}
