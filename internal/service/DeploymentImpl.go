package service

import (
	"fmt"
	"modi/core/result"
	"modi/internal/kind"
	"modi/internal/kind/DomainUtils"
	Model "modi/internal/model/DeploymentModel"
	"modi/pkg/utils"
)

var DeploymentServiceGetter IDeployment

type IDeploymentServiceGetterImpl struct {
}

func (I IDeploymentServiceGetterImpl) GetDeploymentDetailByNsDName(ns string, dname string) *result.ErrorResult {
	dep, err := kind.DeploymentMapInstance.GetDeploymentByName(ns, dname)
	if err != nil {
		return result.Result(nil, fmt.Errorf("GetDeployment: record not found"))
	} else {
		var ret *Model.DeploymentImpl
		ret = Model.New(
			Model.WithName(dep.Name),
			Model.WithNamespace(dep.Namespace),
			Model.WithCreateTime(utils.FormatTime(dep.CreationTimestamp)),
			Model.WithReplicas([3]int32{dep.Status.Replicas, dep.Status.AvailableReplicas, dep.Status.UnavailableReplicas}),
			Model.WithImages(DomainUtils.GetImages(*dep)),
			Model.WithPods(DomainUtils.GetPods(*dep, ns, dname)),
		)
		return result.Result(ret, nil)
	}
}

func (I IDeploymentServiceGetterImpl) IncrReplicas(ns string, dep string, dec bool) *result.ErrorResult {
	isSucceed, err := DomainUtils.IncreaseReplicas(ns, dep, dec)
	return result.Result(isSucceed, err)
}

func (I IDeploymentServiceGetterImpl) GetDeploymentsByNs(ns string) *result.ErrorResult {
	list, err := kind.DeploymentMapInstance.GetDeploymentsByNs(ns)
	if err != nil {
		return result.Result(nil, fmt.Errorf("record not found"))
	} else {
		var ret []*Model.DeploymentImpl
		for _, item := range list {
			ret = append(ret, Model.New(
				Model.WithName(item.Name),
				Model.WithNamespace(item.Namespace),
				Model.WithCreateTime(utils.FormatTime(item.CreationTimestamp)),
				Model.WithReplicas([3]int32{item.Status.Replicas, item.Status.AvailableReplicas, item.Status.UnavailableReplicas}),
				Model.WithImages(DomainUtils.GetImages(*item)),
			))
		}
		return result.Result(ret, nil)
	}
}

func init() {
	DeploymentServiceGetter = &IDeploymentServiceGetterImpl{}
}
