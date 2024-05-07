package DomainUtils

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"modi/client"
	"modi/internal/kind"
	"modi/internal/model/PodModel"
	"sort"
)

func GetPods(dep v1.Deployment, ns string, dname string) []*PodModel.PodImpl {
	var ret []*PodModel.PodImpl

	// 取得所有的 rs
	rsList, err := kind.RsMapInstance.RsByNs(ns)
	if err != nil {
		return nil
	}
	// 根据deployment过滤出rs，然后直接取得labels
	labels, err := GetrslablebydeploymentListWatch(dep, rsList) // 26....labels map[app:prod pod-template-hash:57b8c559dd]
	ret = ListPodsByLabels(ns, labels)

	return ret
}

func ListPodsByLabels(ns string, labels []map[string]string) []*PodModel.PodImpl {
	pods, err := kind.PodMapInstance.ListByLabel(ns, labels)
	if err != nil {
		return nil
	}
	ret := RenderPods(pods)
	return ret
}

func RenderPods(pods []*corev1.Pod) []*PodModel.PodImpl {
	var ret []*PodModel.PodImpl

	sortPods := kind.CoreV1Pods(pods)
	sort.Sort(sortPods)

	for _, item := range sortPods {
		ret = append(ret, PodModel.New(
			PodModel.WithName(item.Name),
			PodModel.WithImages(GetImagesByPod(item.Spec.Containers)),
			PodModel.WithPhase(string(item.Status.Phase)),
			PodModel.WithNodeName(item.Spec.NodeName),
			PodModel.WithCreateTime(item.CreationTimestamp.String()),
			PodModel.WithMessage(GetPodMessage(*item)),
			PodModel.WithIsReady(GetPodIsReady(*item)),
			PodModel.WithNamespace(item.Namespace),
			// todo 在pod中反向取 deployment 名称，待解决
			//PodModel.WithDeploymentName(item.Labels["app"]),
		))
	}
	return ret
}

func GetPodMessage(pod corev1.Pod) string {
	message := ""
	for _, condition := range pod.Status.Conditions {
		if condition.Status != "True" {
			message += condition.Message
		}
	}
	return message
}

// 判断POD 是否就绪
// pod 有一个PodStatus对象，其中包含一个PodCondition 数组。Pod可能通过也可能未通过其中的一些状况测试；
// PodScheduled： Pod 已经被调度到某节点；
// ContainersReady: Pod中所有容器都已就绪；
// Initialized： 所有的init容器 都已经成功启动；
// Ready： Pod可以为请求提供服务，并且应该被添加到对应服务的负载均衡池中；

// 字段名称              描述
// type                  Pod状况的名称    //  PodScheduled， ContainersReady， Initialized， Ready
// status                表明该状况是否适用，可能取值 True， False Unknown
// lastProbeTime			上次探测Pod状况时的时间戳
// LastTransitionTime		Pod上次从一种状态转换到另一种状态时的时间戳
// reason					机器可读的，驼峰编码（UpperCamelCase）的文字，表述上次状况变化的原因；

func GetPodIsReady(pod corev1.Pod) bool {
	// 1. 先判断 Pod中所有容器已经就绪
	for _, condition := range pod.Status.Conditions {
		if condition.Type == "ContainersReady" && condition.Status != "True" {
			return false
		}
	}
	// readinessGates中的所有状况都为True
	for _, rg := range pod.Spec.ReadinessGates {
		for _, condition := range pod.Status.Conditions {
			if condition.Type == rg.ConditionType && condition.Status != "True" {
				return false
			}
		}
	}
	// 满足上述两个条件之后，Pod才会被评估为就绪；
	return true
}

func DeletePod(ns string, podName string) (bool, error) {
	err := client.K8sClient.CoreV1().Pods(ns).Delete(context.Background(), podName, metav1.DeleteOptions{})
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
