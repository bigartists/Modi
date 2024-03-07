package DeploymentModel

import "modi/internal/model/PodModel"

type DeploymentImpl struct {
	Name       string
	Namespace  string
	Replicas   [3]int32
	Images     string
	CreateTime string
	Pods       []*PodModel.PodImpl
}

func New(attrs ...DeploymentAttrFunc) *DeploymentImpl {
	dep := &DeploymentImpl{}
	//attrs.(*DeploymentAttrFuncs).apply(dep)
	DeploymentAttrFuncs(attrs).apply(dep)
	return dep
}
