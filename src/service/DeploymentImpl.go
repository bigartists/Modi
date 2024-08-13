package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"modi/client"
	"modi/src/model/DeploymentModel"
	"modi/src/model/PodModel"
	"modi/src/result"
	"modi/src/utils"
	"sort"
)

var DeploymentServiceGetter IDeployment

type IDeploymentServiceGetterImpl struct {
}

func (I IDeploymentServiceGetterImpl) GetPodContainer(ns string, podName string) *result.ErrorResult {
	podService := PodService{}
	ret := podService.GetPodContainer(ns, podName)
	return result.Result(ret, nil)
}

func (I IDeploymentServiceGetterImpl) GetPodLogs(c *gin.Context, ns string, podname string, cname string) *result.ErrorResult {

	req := client.K8sClient.CoreV1().Pods(ns).GetLogs(podname, &corev1.PodLogOptions{
		//Follow:    true,
		Container: cname,
	})
	podLogs := req.Do(context.Background())
	b, _ := podLogs.Raw()
	//println(string(b))
	return result.Result(string(b), nil)
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

func (I IDeploymentServiceGetterImpl) GetPodDetail(ns string, pod string) *result.ErrorResult {
	podDetail, err := PodMapInstance.GetDetail(ns, pod)
	if err != nil {
		return result.Result(nil, fmt.Errorf("GetPodDetail: record not found"))
	} else {
		//return result.Result(podDetail, nil)
		pods := make([]*corev1.Pod, 0)
		pods = append(pods, podDetail)
		ret := RenderPods(pods)
		return result.Result(ret[0], nil)

	}
}

func (I IDeploymentServiceGetterImpl) GetDeploymentDetailByNsDName(ns string, dname string) *result.ErrorResult {
	dep, err := DeploymentMapInstance.GetDeploymentByName(ns, dname)
	if err != nil {
		return result.Result(nil, fmt.Errorf("GetDeployment: record not found"))
	} else {
		var ret *DeploymentModel.DeploymentImpl
		ret = DeploymentModel.New(
			DeploymentModel.WithName(dep.Name),
			DeploymentModel.WithNamespace(dep.Namespace),
			DeploymentModel.WithCreateTime(utils.FormatTime(dep.CreationTimestamp)),
			DeploymentModel.WithReplicas([3]int32{dep.Status.Replicas, dep.Status.AvailableReplicas, dep.Status.UnavailableReplicas}),
			DeploymentModel.WithImages(GetImages(*dep)),
			DeploymentModel.WithPods(GetPods(*dep, ns, dname)),
			DeploymentModel.WithIsComplete(GetDeploymentIsComplete(dep)),
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
		var ret []*DeploymentModel.DeploymentImpl
		sortList := CoreV1Deployments(list)
		sort.Sort(sortList)

		for _, item := range sortList {
			ret = append(ret, DeploymentModel.New(
				DeploymentModel.WithName(item.Name),
				DeploymentModel.WithNamespace(item.Namespace),
				DeploymentModel.WithCreateTime(utils.FormatTime(item.CreationTimestamp)),
				DeploymentModel.WithReplicas([3]int32{item.Status.Replicas, item.Status.AvailableReplicas, item.Status.UnavailableReplicas}),
				DeploymentModel.WithImages(GetImages(*item)),
				DeploymentModel.WithIsComplete(GetDeploymentIsComplete(item)),
				DeploymentModel.WithMessage(GetDeploymentCondition(item)),
			))
		}
		return result.Result(ret, nil)
	}
}

func init() {
	DeploymentServiceGetter = &IDeploymentServiceGetterImpl{}
}
