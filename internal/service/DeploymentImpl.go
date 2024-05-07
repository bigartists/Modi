package service

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"modi/core/result"
	"modi/internal/kind"
	"modi/internal/kind/DomainUtils"
	Model "modi/internal/model/DeploymentModel"
	"modi/internal/model/PodModel"
	"modi/pkg/utils"
	"sort"
)

var DeploymentServiceGetter IDeployment

type IDeploymentServiceGetterImpl struct {
}

func (I IDeploymentServiceGetterImpl) GetNs() *result.ErrorResult {
	return result.Result(kind.NamespaceMapInstance.GetAllNamespaces(), nil)
}

func (I IDeploymentServiceGetterImpl) DeletePod(ns string, pod string) *result.ErrorResult {
	ret, err := DomainUtils.DeletePod(ns, pod)
	if err != nil {
		return result.Result(nil, err)
	} else {
		return result.Result(ret, nil)
	}
}

func (I IDeploymentServiceGetterImpl) GetPodJson(ns string, pod string) *result.ErrorResult {
	json, err := kind.PodMapInstance.GetDetail(ns, pod)
	if err != nil {
		return result.Result(nil, fmt.Errorf("getPodJson: record not found"))
	} else {
		return result.Result(json, nil)
	}

}

func (I IDeploymentServiceGetterImpl) GetPods(ns string, dname string) *result.ErrorResult {
	var pods []*corev1.Pod
	var err error
	if dname == "" {
		pods, err = kind.PodMapInstance.GetAllPods()
		if err != nil {
			return nil
		}
		podsList := DomainUtils.RenderPods(pods)
		if ns == "" {
			return result.Result(podsList, nil)
		} else {
			var ret []*PodModel.PodImpl
			for _, item := range podsList {
				if item.Namespace == ns {
					ret = append(ret, item)
				}
			}
			return result.Result(ret, nil)
		}
	} else {
		dep, err := kind.DeploymentMapInstance.GetDeploymentByName(ns, dname)
		if err != nil {
			return result.Result(nil, fmt.Errorf("GetDeployment: record not found"))
		}
		pods := DomainUtils.GetPods(*dep, ns, dname)
		return result.Result(pods, nil)
	}
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
	var list []*v1.Deployment
	var err error
	if ns == "" {
		list, err = kind.DeploymentMapInstance.GetAllDeployment()
	} else {
		list, err = kind.DeploymentMapInstance.GetDeploymentsByNs(ns)
	}

	if err != nil {
		return result.Result(nil, fmt.Errorf("record not found"))
	} else {
		var ret []*Model.DeploymentImpl
		sortList := kind.CoreV1Deployments(list)
		sort.Sort(sortList)

		for _, item := range sortList {
			ret = append(ret, Model.New(
				Model.WithName(item.Name),
				Model.WithNamespace(item.Namespace),
				Model.WithCreateTime(utils.FormatTime(item.CreationTimestamp)),
				Model.WithReplicas([3]int32{item.Status.Replicas, item.Status.AvailableReplicas, item.Status.UnavailableReplicas}),
				Model.WithImages(DomainUtils.GetImages(*item)),
				Model.WithIsComplete(DomainUtils.GetDeploymentIsComplete(item)),
				Model.WithMessage(DomainUtils.GetDeploymentCondition(item)),
			))
		}
		return result.Result(ret, nil)
	}
}

func init() {
	DeploymentServiceGetter = &IDeploymentServiceGetterImpl{}
}
