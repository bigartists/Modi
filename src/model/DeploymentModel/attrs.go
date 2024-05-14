package DeploymentModel

import (
	"modi/src/model/PodModel"
)

type DeploymentAttrFunc func(*DeploymentImpl)
type DeploymentAttrFuncs []DeploymentAttrFunc

func WithName(name string) DeploymentAttrFunc {
	return func(dep *DeploymentImpl) {
		dep.Name = name
	}
}

func WithNamespace(ns string) DeploymentAttrFunc {
	return func(dep *DeploymentImpl) {
		dep.Namespace = ns
	}
}

func WithReplicas(replicas [3]int32) DeploymentAttrFunc {
	return func(dep *DeploymentImpl) {
		dep.Replicas = replicas
	}
}

func WithIsComplete(is bool) DeploymentAttrFunc {
	return func(dep *DeploymentImpl) {
		dep.IsComplete = is
	}
}

func WithMessage(msg string) DeploymentAttrFunc {
	return func(dep *DeploymentImpl) {
		dep.Message = msg
	}
}

func WithImages(images string) DeploymentAttrFunc {
	return func(dep *DeploymentImpl) {
		dep.Images = images
	}
}

func WithCreateTime(createTime string) DeploymentAttrFunc {
	return func(dep *DeploymentImpl) {
		dep.CreateTime = createTime
	}
}

func WithPods(pods []*PodModel.PodImpl) DeploymentAttrFunc {
	return func(dep *DeploymentImpl) {
		dep.Pods = pods
	}
}

func (this DeploymentAttrFuncs) apply(d *DeploymentImpl) {
	for _, f := range this {
		f(d)
	}
}
