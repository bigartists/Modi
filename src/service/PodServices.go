package service

import (
	"github.com/bigartists/Modi/src/model"
	"github.com/bigartists/Modi/src/repo"
)

type PodService struct {
	podRepo *repo.PodRepo
}

func ProviderPodService(podRepo *repo.PodRepo) *PodService {
	return &PodService{podRepo: podRepo}
}

func (this *PodService) GetPodContainer(ns, podname string) []*model.ContainerModel {
	ret := make([]*model.ContainerModel, 0)
	pod := this.podRepo.Get(ns, podname)
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
