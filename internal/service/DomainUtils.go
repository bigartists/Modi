package service

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"modi/client"
	"modi/internal/model/PodModel"
	"sort"
)

func GetPods(dep v1.Deployment, ns string, dname string) []*PodModel.PodImpl {
	var ret []*PodModel.PodImpl

	// 取得所有的 rs
	rsList, err := RsMapInstance.RsByNs(ns)
	if err != nil {
		return nil
	}
	// 根据deployment过滤出rs，然后直接取得labels
	labels, err := GetrslablebydeploymentListWatch(dep, rsList) // 26....labels map[app:prod pod-template-hash:57b8c559dd]
	ret = ListPodsByLabels(ns, labels)

	return ret
}

func ListPodsByLabels(ns string, labels []map[string]string) []*PodModel.PodImpl {
	pods, err := PodMapInstance.ListByLabel(ns, labels)
	if err != nil {
		return nil
	}
	ret := RenderPods(pods)
	return ret
}

func RenderPods(pods []*corev1.Pod) []*PodModel.PodImpl {
	var ret []*PodModel.PodImpl

	sortPods := CoreV1Pods(pods)
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

func GetLabels(m map[string]string) string {
	labels := ""
	// aa=xxx, xxx=xx
	for k, v := range m {
		if labels != "" {
			labels += ","
		}
		labels += fmt.Sprintf("%s=%s", k, v)
	}
	return labels
}

func GetImagesByPod(container []corev1.Container) string {
	images := container[0].Image
	if len(container) > 1 {
		images += fmt.Sprintf("+其他%d个镜像", len(container)-1)
	}
	return images
}

func GetImages(dep v1.Deployment) string {
	return GetImagesByPod(dep.Spec.Template.Spec.Containers)
}

func IncreaseReplicas(ns string, dep string, dec bool) (bool, error) {
	ctx := context.Background()
	scale, err := client.K8sClient.AppsV1().Deployments(ns).GetScale(ctx, dep, metav1.GetOptions{})
	if err != nil {
		return false, err
	}
	if dec { // dec == true 减少副本数
		scale.Spec.Replicas--
	} else { // dec == false 增加副本数
		scale.Spec.Replicas++
	}
	_, err = client.K8sClient.AppsV1().Deployments(ns).UpdateScale(ctx, dep, scale, metav1.UpdateOptions{})
	if err != nil {
		return false, err
	}
	return true, nil
}

// 取出跟dep match的所有 rs，而不是最新的。 解决 2/1/1的问题；
func GetrslablebydeploymentListWatch(dep v1.Deployment, rslist []*v1.ReplicaSet) ([]map[string]string, error) {
	ret := make([]map[string]string, 0)
	for _, item := range rslist {
		if IsRsFromDep(dep, *item) {
			s, err := metav1.LabelSelectorAsMap(item.Spec.Selector)
			if err != nil {
				return nil, err
			}
			ret = append(ret, s)
		}
	}
	return ret, nil
}

// 判断当前的rs 是否是最新的；因此这个方法返回的 rs都是最新的rs；所以只会显示一个，而前面报错的就没法显示了；就导致访问dep显示 2/1/1，但pod只有报错的那个；

func IsCurrentRsByDep(dep v1.Deployment, set v1.ReplicaSet) bool {
	// 下面这一步是 判断是否是最新
	if set.ObjectMeta.Annotations["deployment.kubernetes.io/revision"] != dep.ObjectMeta.Annotations["deployment.kubernetes.io/revision"] {
		return false
	}
	return IsRsFromDep(dep, set)
}

// 判断 rs 是否属于某个dep，而不是判断它是最新的。 因此，得调用这个，才能解决 2/1/1的问题；
func IsRsFromDep(dep v1.Deployment, set v1.ReplicaSet) bool {
	for _, ref := range set.OwnerReferences {
		if ref.Kind == "Deployment" && ref.Name == dep.Name {
			return true
		}
	}
	return false
}

// 判断deployment是否完成

func GetDeploymentIsComplete(dep *v1.Deployment) bool {
	return dep.Status.Replicas == dep.Status.AvailableReplicas
}

// 获取deployment失败状态

func GetDeploymentCondition(dep *v1.Deployment) string {
	for _, item := range dep.Status.Conditions {
		if string(item.Type) == "Available" && string(item.Status) != "True" {
			return item.Message
		}
	}
	return ""
}
