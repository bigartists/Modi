package kind

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"log"
	"sync"
)

type PodMapStruct struct {
	data sync.Map
}

func (this *PodMapStruct) Add(pod *corev1.Pod) {
	if list, ok := this.data.Load(pod.Namespace); ok {
		list = append(list.([]*corev1.Pod), pod)
		this.data.Store(pod.Namespace, list)
	} else {
		this.data.Store(pod.Namespace, []*corev1.Pod{pod})
	}
}

func (this *PodMapStruct) Update(pod *corev1.Pod) error {
	if list, ok := this.data.Load(pod.Namespace); ok {
		for i, range_pod := range list.([]*corev1.Pod) {
			if range_pod.Name == pod.Name {
				list.([]*corev1.Pod)[i] = pod
			}
		}
		return nil
	}
	return fmt.Errorf("Pod-%s not found", pod.Name)
}

func (this *PodMapStruct) Delete(pod *corev1.Pod) {
	if list, ok := this.data.Load(pod.Namespace); ok {
		for i, range_pod := range list.([]*corev1.Pod) {
			if range_pod.Name == pod.Name {
				newList := append(list.([]*corev1.Pod)[:i], list.([]*corev1.Pod)[i+1:]...)
				this.data.Store(pod.Namespace, newList)
				break
			}
		}
	}
}

//func (this *PodMapStruct) ListByLabel(ns string, labels map[string]string) ([]*corev1.Pod, error) {
//	ret := make([]*corev1.Pod, 0)
//	if list, ok := this.data.Load(ns); ok {
//		for _, pod := range list.([]*corev1.Pod) {
//			if reflect.DeepEqual(pod.Labels, labels) { //标签完全匹配
//				ret = append(ret, pod)
//			}
//		}
//		return ret, nil
//	}
//	return nil, fmt.Errorf("ListByLabel record not found")
//}

func (this *PodMapStruct) ListByLabel(ns string, labels map[string]string) ([]*corev1.Pod, error) {
	fmt.Println("ns=", ns, "labels=", labels)
	ret := make([]*corev1.Pod, 0)
	if list, ok := this.data.Load(ns); ok {

		for _, pod := range list.([]*corev1.Pod) {
			// fmt.Println("pod.name=", pod.Name, pod.Labels)
			//if reflect.DeepEqual(pod.Labels, labels) { //标签完全匹配,因为pod中有istio的信息，所以这个有个坑；
			//	ret = append(ret, pod)
			//}
			//1
			//ns= groot labels= map[app:prod pod-template-hash:57b8c559dd]
			//pod.name= reviewapi-6b9d748877-6lwwp map[app:reviews pod-template-hash:6b9d748877 security.istio.io/tlsMode:istio service.istio.io/canonical-name:reviews service.istio.io/canonical-revision:latest]
			//pod.name= prodapi-57b8c559dd-5vsdj map[app:prod pod-template-hash:57b8c559dd security.istio.io/tlsMode:istio service.istio.io/canonical-name:prod service.istio.io/canonical-revision:latest]
			//pod.name= prodapi-57b8c559dd-pkbwm map[app:prod pod-template-hash:57b8c559dd security.istio.io/tlsMode:istio service.istio.io/canonical-name:prod service.istio.io/canonical-revision:latest]

			isSubset := true
			for key, value := range labels {
				if podValue, found := pod.Labels[key]; !found || podValue != value {
					isSubset = false
					break
				}
			}
			if isSubset {
				ret = append(ret, pod)
			}
		}
		return ret, nil
	}
	return nil, fmt.Errorf("pods not found ")
}

var PodMapInstance *PodMapStruct

type PodHandler struct{}

func (p PodHandler) OnAdd(obj interface{}, isInInitialList bool) {
	PodMapInstance.Add(obj.(*corev1.Pod))
}

func (p PodHandler) OnUpdate(oldObj, newObj interface{}) {
	err := PodMapInstance.Update(newObj.(*corev1.Pod))
	if err != nil {
		log.Println(err)
	}
}

func (p PodHandler) OnDelete(obj interface{}) {
	if d, ok := obj.(*corev1.Pod); ok {
		PodMapInstance.Delete(d)
	}
}

func init() {
	PodMapInstance = &PodMapStruct{}
}
