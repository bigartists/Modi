package service

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

func (this *EventMapStruct) Store(key string, event *v1.Event) {
	this.data.Store(key, event)
}

func (this *EventMapStruct) Delete(key string) {
	this.data.Delete(key)
}

func (this *EventMapStruct) GetMessage(ns string, kind string, name string) string {
	key := fmt.Sprintf("%s_%s_%s", ns, kind, name)
	if v, ok := this.data.Load(key); ok {
		return v.(*v1.Event).Message
	}
	return ""
}

func init() {
	EventMapInstance = &EventMapStruct{}
}
