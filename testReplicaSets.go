package main

import (
	"context"
	"fmt"
	"github.com/bigartists/Modi/client"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func main() {
	// 获取dep
	dep, _ := client.K8sClient.AppsV1().Deployments("groot").Get(context.Background(), "prodapi", v1.GetOptions{})
	fmt.Println(dep.Namespace, dep.Name) // groot prodapi

	fmt.Println(dep.Spec.Selector) // &LabelSelector{MatchLabels:map[string]string{app: prod,},MatchExpressions:[]LabelSelectorRequirement{},}

	selector, err := v1.LabelSelectorAsSelector(dep.Spec.Selector)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(selector.String()) // app=prod

	list := v1.ListOptions{LabelSelector: selector.String()}

	rs, _ := client.K8sClient.AppsV1().ReplicaSets("groot").List(context.Background(), list)

	for _, v := range rs.Items {
		fmt.Println(v.Name) // prodapi-57b8c559dd
		fmt.Println(IsCurrentRsByDep(dep, v))
	}
}

func IsCurrentRsByDep(dep *appsv1.Deployment, set appsv1.ReplicaSet) bool {
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
