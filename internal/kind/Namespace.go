package kind

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

type NamespaceMap struct {
	data sync.Map
}

func (this *NamespaceMap) Add(namespace *corev1.Namespace) {
	this.data.Store(namespace.Name, namespace)
}

func (this *NamespaceMap) Update(namespace *corev1.Namespace) error {
	_, exists := this.data.Load(namespace.Name)
	if exists {
		this.data.Store(namespace.Name, namespace)
		return nil
	}
	return fmt.Errorf("namespace-%s not found", namespace.Name)
}

func (this *NamespaceMap) Delete(namespace *corev1.Namespace) {
	this.data.Delete(namespace.Name)
}

func (this *NamespaceMap) GetNamespaceByName(name string) (*corev1.Namespace, error) {
	if ns, ok := this.data.Load(name); ok {
		return ns.(*corev1.Namespace), nil
	}
	return nil, fmt.Errorf("GetNamespace: record not found")
}

func (this *NamespaceMap) GetAllNamespaces() []string {
	items := convertToMapItems(this.data)
	sort.Sort(items)
	namespaces := make([]string, len(items))
	for index, item := range items {
		namespaces[index] = item.key
	}
	return namespaces
}

type NamespaceHandler struct {
}

var NamespaceMapInstance *NamespaceMap

func init() {
	NamespaceMapInstance = &NamespaceMap{}
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
