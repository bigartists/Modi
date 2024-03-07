package DomainUtils

import (
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"modi/internal/kind"
	"modi/internal/model/PodModel"
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

	ret = ListPodsByLabels(ns, labels) // todo 问题大概出在这里；

	return ret
}

func ListPodsByLabels(ns string, labels map[string]string) []*PodModel.PodImpl {
	var ret []*PodModel.PodImpl
	pods, err := kind.PodMapInstance.ListByLabel(ns, labels)
	if err != nil {
		return nil
	}
	for _, item := range pods {
		ret = append(ret, PodModel.New(
			PodModel.WithName(item.Name),
			PodModel.WithImages(GetImagesByPod(item.Spec.Containers)),
			PodModel.WithPhase(string(item.Status.Phase)),
			PodModel.WithNodeName(item.Spec.NodeName),
			PodModel.WithCreateTime(item.CreationTimestamp.String()),
			PodModel.WithMessage(GetPodMessage(*item)),
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
