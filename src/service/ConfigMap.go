package service

import (
	"github.com/bigartists/Modi/src/helpers"
	"github.com/bigartists/Modi/src/result"
	corev1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

type cm struct {
	cmdata *corev1.ConfigMap
	md5    string //对cm的data进行md5存储，防止过度更新
}
type CoreV1ConfigMap []*cm

func (this CoreV1ConfigMap) Len() int {
	return len(this)
}
func (this CoreV1ConfigMap) Less(i, j int) bool {
	//根据时间排序    倒排序
	return this[i].cmdata.CreationTimestamp.Time.After(this[j].cmdata.CreationTimestamp.Time)
}
func (this CoreV1ConfigMap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// SecretMap
type ConfigMapStruct struct {
	data sync.Map // [ns string] []*v1.Secret
}

func newcm(c *corev1.ConfigMap) *cm {
	return &cm{
		cmdata: c, //原始对象
		md5:    helpers.Md5Data(c.Data),
	}
}

func (this *ConfigMapStruct) Get(ns string, name string) *corev1.ConfigMap {
	if items, ok := this.data.Load(ns); ok {
		for _, item := range items.([]*cm) {
			if item.cmdata.Name == name {
				return item.cmdata
			}
		}
	}
	return nil
}
func (this *ConfigMapStruct) Add(item *corev1.ConfigMap) {
	if list, ok := this.data.Load(item.Namespace); ok {
		list = append(list.([]*cm), newcm(item))
		this.data.Store(item.Namespace, list)
	} else {
		this.data.Store(item.Namespace, []*cm{newcm(item)})
	}
}

// 返回值 是true 或false . true代表有值更新了， 否则返回false
func (this *ConfigMapStruct) Update(item *corev1.ConfigMap) bool {
	if list, ok := this.data.Load(item.Namespace); ok {
		for i, range_item := range list.([]*cm) {
			//这里做判断，如果没变化就不做 更新
			if range_item.cmdata.Name == item.Name && !helpers.CmIsEq(range_item.cmdata.Data, item.Data) {
				list.([]*cm)[i] = newcm(item)
				return true //代表有值更新了
			}
		}
	}
	return false
}
func (this *ConfigMapStruct) Delete(svc *corev1.ConfigMap) {
	if list, ok := this.data.Load(svc.Namespace); ok {
		for i, range_item := range list.([]*cm) {
			if range_item.cmdata.Name == svc.Name {
				newList := append(list.([]*cm)[:i], list.([]*cm)[i+1:]...)
				this.data.Store(svc.Namespace, newList)
				break
			}
		}
	}
}
func (this *ConfigMapStruct) ListAll(ns string) *result.ErrorResult {
	ret := []*corev1.ConfigMap{}
	if list, ok := this.data.Load(ns); ok {
		newList := list.([]*cm)
		sort.Sort(CoreV1ConfigMap(newList)) //  按时间倒排序
		for _, cm := range newList {
			ret = append(ret, cm.cmdata)
		}
	}
	//return ret, nil //返回空列表
	return result.Result(ret, nil)
}

var ConfigMapInstance *ConfigMapStruct

func init() {
	ConfigMapInstance = &ConfigMapStruct{}
}
