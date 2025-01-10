package repo

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"sync"
)

// EventMap集合 用来保存事件，只保存最新的一条

type EventRepo struct {
	data sync.Map // [key string] *v1.Event
	// key=>namespace+"_"+kind+"_"+name   确保key的唯一性；
}

func ProviderEventRepo() *EventRepo {
	return &EventRepo{data: sync.Map{}}
}

func (this *EventRepo) storeData(obj interface{}, isDelete bool) {
	if event, ok := obj.(*corev1.Event); ok {
		key := fmt.Sprintf("%s_%s_%s", event.Namespace, event.InvolvedObject.Kind, event.InvolvedObject.Name)
		//if event.Namespace == "groot" {
		//	fmt.Println(".....event.Message.......", "key=", key, event.Message)
		//}
		//if key == "groot_Pod_ng1-5cb776c5f-wq4ld" {
		//	fmt.Println(".....event.Message.......", "key=", key, "eventType=", event.Type, "value=", event.Message)
		//}

		if !isDelete {
			this.data.Store(key, event)
		} else {
			this.data.Delete(key)
		}
	}
}

func (this *EventRepo) OnAdd(obj interface{}, isInInitialList bool) {
	this.storeData(obj, false)
}

func (this *EventRepo) OnUpdate(oldObj, newObj interface{}) {
	this.storeData(newObj, false)
}

func (this *EventRepo) OnDelete(obj interface{}) {
	this.storeData(obj, true)
}

func (this *EventRepo) Store(key string, event *v1.Event) {
	this.data.Store(key, event)
}

func (this *EventRepo) Delete(key string) {
	this.data.Delete(key)
}

func (this *EventRepo) GetMessage(ns string, kind string, name string) string {
	key := fmt.Sprintf("%s_%s_%s", ns, kind, name)
	if v, ok := this.data.Load(key); ok {
		return v.(*v1.Event).Message
	}
	return ""
}
