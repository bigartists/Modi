package DeploymentModel

import (
	"modi/pkg/model/PodModel"
)

type DeploymentImpl struct {
	Name       string
	Namespace  string
	Replicas   [3]int32
	Images     string
	CreateTime string
	Pods       []*PodModel.PodImpl
	IsComplete bool // 是否完成
	Message    string
}

func New(attrs ...DeploymentAttrFunc) *DeploymentImpl {
	dep := &DeploymentImpl{}
	//attrs.(*DeploymentAttrFuncs).apply(dep)
	DeploymentAttrFuncs(attrs).apply(dep)
	return dep
}
