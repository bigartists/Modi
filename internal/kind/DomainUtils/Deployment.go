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

func GetrslablebydeploymentListWatch(dep v1.Deployment, rslist []*v1.ReplicaSet) (map[string]string, error) {
	for _, item := range rslist {
		if IsCurrentRsByDep(dep, *item) {
			s, err := metav1.LabelSelectorAsMap(item.Spec.Selector)
			if err != nil {
				return nil, err
			}
			return s, nil
		}
	}
	return nil, nil
}

func IsCurrentRsByDep(dep v1.Deployment, set v1.ReplicaSet) bool {
	if set.ObjectMeta.Annotations["deployment.kubernetes.io/revision"] != dep.ObjectMeta.Annotations["deployment.kubernetes.io/revision"] {
		return false
	}
	for _, ref := range set.OwnerReferences {
		if ref.Kind == "Deployment" && ref.Name == dep.Name {
			return true
		}
	}
	return false
}
