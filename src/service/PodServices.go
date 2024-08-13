package service

import "modi/src/model"

type PodService struct {
}

func (this *PodService) GetPodContainer(ns, podname string) []*model.ContainerModel {
	ret := make([]*model.ContainerModel, 0)
	pod := PodMapInstance.Get(ns, podname)
	if pod != nil {
		for _, c := range pod.Spec.Containers {
			ret = append(ret, &model.ContainerModel{
				Name:  c.Name,
				Image: c.Image,
				//State:c.
			})
		}
	}
	return ret
}
