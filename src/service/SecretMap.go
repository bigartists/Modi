package service

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

type CoreV1Secret []*corev1.Secret

func (this CoreV1Secret) Len() int {
	return len(this)
}
func (this CoreV1Secret) Less(i, j int) bool {
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func (this CoreV1Secret) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// SecretMap
type SecretMapStruct struct {
	data sync.Map // [ns string] []*v1.Secret
}

func (this *SecretMapStruct) Get(ns string, name string) *corev1.Secret {
	if items, ok := this.data.Load(ns); ok {
		for _, item := range items.([]*corev1.Secret) {
			if item.Name == name {
				return item
			}
		}
	}
	return nil
}
func (this *SecretMapStruct) Add(item *corev1.Secret) {
	if list, ok := this.data.Load(item.Namespace); ok {
		list = append(list.([]*corev1.Secret), item)
		this.data.Store(item.Namespace, list)
	} else {
		this.data.Store(item.Namespace, []*corev1.Secret{item})
	}
}
func (this *SecretMapStruct) Update(item *corev1.Secret) error {
	if list, ok := this.data.Load(item.Namespace); ok {
		for i, range_item := range list.([]*corev1.Secret) {
			if range_item.Name == item.Name {
				list.([]*corev1.Secret)[i] = item
			}
		}
		return nil
	}
	return fmt.Errorf("Secret-%s not found", item.Name)
}
func (this *SecretMapStruct) Delete(svc *corev1.Secret) {
	if list, ok := this.data.Load(svc.Namespace); ok {
		for i, range_item := range list.([]*corev1.Secret) {
			if range_item.Name == svc.Name {
				newList := append(list.([]*corev1.Secret)[:i], list.([]*corev1.Secret)[i+1:]...)
				this.data.Store(svc.Namespace, newList)
				break
			}
		}
	}
}
func (this *SecretMapStruct) ListAllByNs(ns string) []*corev1.Secret {
	if list, ok := this.data.Load(ns); ok {
		newList := list.([]*corev1.Secret)
		sort.Sort(CoreV1Secret(newList)) //  按时间倒排序

		return newList
	}
	return []*corev1.Secret{} //返回空列表
}

func (this *SecretMapStruct) ListAll() []*corev1.Secret {
	var ret []*corev1.Secret
	this.data.Range(func(key, value interface{}) bool {
		for _, item := range value.([]*corev1.Secret) {
			ret = append(ret, item)
		}
		return true
	})
	sort.Sort(CoreV1Secret(ret)) //  按时间倒排序
	return ret

}

var SecretMapInstance *SecretMapStruct

func init() {
	SecretMapInstance = &SecretMapStruct{}
}
