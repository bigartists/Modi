package service

import (
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sync"
)

type MapItems []*MapItem

type MapItem struct {
	key   string
	value interface{}
}

// 把sync.map 转为自定义切片
func convertToMapItems(m sync.Map) MapItems {
	items := make(MapItems, 0)
	m.Range(func(key, value interface{}) bool {
		items = append(items, &MapItem{key: key.(string), value: value})
		return true
	})
	return items
}

func (this MapItems) Len() int {
	return len(this)
}

func (this MapItems) Less(i, j int) bool {
	return this[i].key < this[j].key
}

func (this MapItems) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type CoreV1Pods []*corev1.Pod

func (this CoreV1Pods) Len() int {
	return len(this)
}

func (this CoreV1Pods) Less(i, j int) bool {
	return this[i].CreationTimestamp.Time.Before(this[j].CreationTimestamp.Time)
}

func (this CoreV1Pods) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type CoreV1Deployments []*v1.Deployment

func (this CoreV1Deployments) Len() int {
	return len(this)
}

func (this CoreV1Deployments) Less(i, j int) bool {
	// 根据创建时间排序，如果有两个或者三个创建时间一样的话，那么他们内部再按 name 的长度进行排序
	if this[i].CreationTimestamp.Time.Equal(this[j].CreationTimestamp.Time) {
		return len(this[i].Name) < len(this[j].Name)
	}
	return this[i].CreationTimestamp.Time.Before(this[j].CreationTimestamp.Time)
}

func (this CoreV1Deployments) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
