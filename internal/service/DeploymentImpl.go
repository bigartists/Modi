package service

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"modi/internal/result"

	Model "modi/internal/model/DeploymentModel"
	"modi/internal/model/PodModel"
	"modi/pkg/utils"
	"sort"
)

var DeploymentServiceGetter IDeployment

type IDeploymentServiceGetterImpl struct {
}

func (I IDeploymentServiceGetterImpl) GetNs() *result.ErrorResult {
	return result.Result(NamespaceMapInstance.GetAllNamespaces(), nil)
}

func (I IDeploymentServiceGetterImpl) DeletePod(ns string, pod string) *result.ErrorResult {
	ret, err := DeletePod(ns, pod)
	if err != nil {
		return result.Result(nil, err)
	} else {
		return result.Result(ret, nil)
	}
}

func (I IDeploymentServiceGetterImpl) GetPodJson(ns string, pod string) *result.ErrorResult {
	json, err := PodMapInstance.GetDetail(ns, pod)
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
		pods, err = PodMapInstance.GetAllPods()
		if err != nil {
			return nil
		}
		podsList := RenderPods(pods)
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
		dep, err := DeploymentMapInstance.GetDeploymentByName(ns, dname)
		if err != nil {
			return result.Result(nil, fmt.Errorf("GetDeployment: record not found"))
		}
		pods := GetPods(*dep, ns, dname)
		return result.Result(pods, nil)
	}
}

func (I IDeploymentServiceGetterImpl) GetDeploymentDetailByNsDName(ns string, dname string) *result.ErrorResult {
	dep, err := DeploymentMapInstance.GetDeploymentByName(ns, dname)
	if err != nil {
		return result.Result(nil, fmt.Errorf("GetDeployment: record not found"))
	} else {
		var ret *Model.DeploymentImpl
		ret = Model.New(
			Model.WithName(dep.Name),
			Model.WithNamespace(dep.Namespace),
			Model.WithCreateTime(utils.FormatTime(dep.CreationTimestamp)),
			Model.WithReplicas([3]int32{dep.Status.Replicas, dep.Status.AvailableReplicas, dep.Status.UnavailableReplicas}),
			Model.WithImages(GetImages(*dep)),
			Model.WithPods(GetPods(*dep, ns, dname)),
		)
		return result.Result(ret, nil)
	}
}

func (I IDeploymentServiceGetterImpl) IncrReplicas(ns string, dep string, dec bool) *result.ErrorResult {
	isSucceed, err := IncreaseReplicas(ns, dep, dec)
	return result.Result(isSucceed, err)
}

func (I IDeploymentServiceGetterImpl) GetDeploymentsByNs(ns string) *result.ErrorResult {
	var list []*v1.Deployment
	var err error
	if ns == "" {
		list, err = DeploymentMapInstance.GetAllDeployment()
	} else {
		list, err = DeploymentMapInstance.GetDeploymentsByNs(ns)
	}

	if err != nil {
		return result.Result(nil, fmt.Errorf("record not found"))
	} else {
		var ret []*Model.DeploymentImpl
		sortList := CoreV1Deployments(list)
		sort.Sort(sortList)

		for _, item := range sortList {
			ret = append(ret, Model.New(
				Model.WithName(item.Name),
				Model.WithNamespace(item.Namespace),
				Model.WithCreateTime(utils.FormatTime(item.CreationTimestamp)),
				Model.WithReplicas([3]int32{item.Status.Replicas, item.Status.AvailableReplicas, item.Status.UnavailableReplicas}),
				Model.WithImages(GetImages(*item)),
				Model.WithIsComplete(GetDeploymentIsComplete(item)),
				Model.WithMessage(GetDeploymentCondition(item)),
			))
		}
		return result.Result(ret, nil)
	}
}

func init() {
	DeploymentServiceGetter = &IDeploymentServiceGetterImpl{}
}
