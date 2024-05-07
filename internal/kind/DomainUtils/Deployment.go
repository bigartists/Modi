package DomainUtils

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"modi/client"
)

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
