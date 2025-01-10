package repo

import (
	"fmt"
	"github.com/bigartists/Modi/src/utils"
	corev1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

type NamespaceRepo struct {
	data sync.Map
}

func ProviderNamespaceRepo() *NamespaceRepo {
	return &NamespaceRepo{
		data: sync.Map{},
	}
}

func (this *NamespaceRepo) OnAdd(obj interface{}, isInInitialList bool) {
	this.Add(obj.(*corev1.Namespace))
}

func (this *NamespaceRepo) OnUpdate(oldObj, newObj interface{}) {
	err := this.Update(newObj.(*corev1.Namespace))
	if err != nil {
		return
	}
}

func (this *NamespaceRepo) OnDelete(obj interface{}) {
	namespace := obj.(*corev1.Namespace)
	this.Delete(namespace)
}

func (this *NamespaceRepo) Add(namespace *corev1.Namespace) {
	this.data.Store(namespace.Name, namespace)
}

func (this *NamespaceRepo) Update(namespace *corev1.Namespace) error {
	_, exists := this.data.Load(namespace.Name)
	if exists {
		this.data.Store(namespace.Name, namespace)
		return nil
	}
	return fmt.Errorf("namespace-%s not found", namespace.Name)
}

func (this *NamespaceRepo) Delete(namespace *corev1.Namespace) {
	this.data.Delete(namespace.Name)
}

func (this *NamespaceRepo) GetNamespaceByName(name string) (*corev1.Namespace, error) {
	if ns, ok := this.data.Load(name); ok {
		return ns.(*corev1.Namespace), nil
	}
	return nil, fmt.Errorf("GetNamespace: record not found")
}

func (this *NamespaceRepo) GetAllNamespaces() []string {
	items := utils.ConvertToMapItems(this.data)
	sort.Sort(items)
	namespaces := make([]string, len(items))
	for index, item := range items {
		namespaces[index] = item.Key
	}
	return namespaces
}
