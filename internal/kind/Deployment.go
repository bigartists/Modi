package kind

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"sync"
)

type DeploymentMap struct {
	data sync.Map
}

func (this *DeploymentMap) Add(deployment *v1.Deployment) {
	ns := deployment.Namespace
	if list, ok := this.data.Load(ns); ok {
		list = append(list.([]*v1.Deployment), deployment)
		this.data.Store(ns, list)
	} else {
		this.data.Store(ns, []*v1.Deployment{deployment})
	}
}

func (this *DeploymentMap) Update(deployment *v1.Deployment) error {
	ns := deployment.Namespace
	if list, ok := this.data.Load(ns); ok {
		for item, rangeDep := range list.([]*v1.Deployment) {
			if rangeDep.Name == deployment.Name {
				list.([]*v1.Deployment)[item] = deployment
			}
		}
		return nil
	}
	return fmt.Errorf("deployment-%s not found", deployment.Name)
}

func (this *DeploymentMap) Delete(deployment *v1.Deployment) {
	ns := deployment.Namespace
	if list, ok := this.data.Load(ns); ok {
		for item, rangeDep := range list.([]*v1.Deployment) {
			if rangeDep.Name == deployment.Name {
				// list.([]*v1.Deployment)[:item]：这是原切片的一个切片，包含了从索引0到item-1的所有元素
				// list.([]*v1.Deployment)[item+1:]：这是原切片的另一个切片，包含了从索引item+1到切片末尾的所有元素
				// append(list.([]*v1.Deployment)[:item], list.([]*v1.Deployment)[item+1:]...)：这是将两个切片合并成一个新的切片
				// 第一个参数是新切片的开始部分
				// 第二个参数是...，它是一个切片展开操作符，用于将第二个切片中的所有元素作为独立参数传递给append函数；
				newList := append(list.([]*v1.Deployment)[:item], list.([]*v1.Deployment)[item+1:]...)
				this.data.Store(ns, newList)
				break
			}
		}
	}
}

func (this *DeploymentMap) GetDeploymentsByNs(ns string) ([]*v1.Deployment, error) {
	if list, ok := this.data.Load(ns); ok {
		//return list.([]*v1.Deployment), nil
		return list.([]*v1.Deployment), nil
	} else {
		//return nil, fmt.Errorf("record not found")
		return nil, fmt.Errorf("record not found")
	}
}

func (this *DeploymentMap) GetAllDeployment() ([]*v1.Deployment, error) {
	var lists []*v1.Deployment
	this.data.Range(func(key, value any) bool {
		lists = append(lists, value.([]*v1.Deployment)...)
		return true
	})

	if len(lists) == 0 {
		return nil, fmt.Errorf("no deployment found")
	}

	return lists, nil
}

func (this *DeploymentMap) GetDeploymentByName(ns string, name string) (*v1.Deployment, error) {
	if list, ok := this.data.Load(ns); ok {
		for _, dep := range list.([]*v1.Deployment) {
			if dep.Name == name {
				return dep, nil
			}
		}
	}
	return nil, fmt.Errorf("GetDeployment: record not found")
}

type DeploymentHandler struct {
}

var DeploymentMapInstance *DeploymentMap

func init() {
	DeploymentMapInstance = &DeploymentMap{}
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
