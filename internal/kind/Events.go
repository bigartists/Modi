package kind

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"sync"
)

// EventMap集合 用来保存事件，只保存最新的一条

type EventMapStruct struct {
	data sync.Map // [key string] *v1.Event
	// key=>namespace+"_"+kind+"_"+name   确保key的唯一性；
}

var EventMapInstance *EventMapStruct

type EventHandler struct {
}

func (this *EventHandler) storeData(obj interface{}, isDelete bool) {
	if event, ok := obj.(*v1.Event); ok {
		key := fmt.Sprintf("%s_%s_%s", event.Namespace, event.InvolvedObject.Kind, event.InvolvedObject.Name)
		//if event.Namespace == "groot" {
		//	fmt.Println(".....event.Message.......", "key=", key, event.Message)
		//}
		//if key == "groot_Pod_ng1-5cb776c5f-wq4ld" {
		//	fmt.Println(".....event.Message.......", "key=", key, "eventType=", event.Type, "value=", event.Message)
		//}

		if !isDelete {
			EventMapInstance.data.Store(key, event)
		} else {
			EventMapInstance.data.Delete(key)
		}
	}
}

func (this *EventMapStruct) GetMessage(ns string, kind string, name string) string {
	key := fmt.Sprintf("%s_%s_%s", ns, kind, name)
	if v, ok := this.data.Load(key); ok {
		return v.(*v1.Event).Message
	}
	return ""
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

func init() {
	EventMapInstance = &EventMapStruct{}
}
