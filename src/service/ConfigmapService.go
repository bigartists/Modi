package service

import (
	"github.com/bigartists/Modi/src/repo"
	"github.com/bigartists/Modi/src/result"
)

type ConfigmapService struct {
	configmapRepo *repo.ConfigMapRepo
}

func ProviderConfigmapService(configmapRepo *repo.ConfigMapRepo) *ConfigmapService {
	return &ConfigmapService{configmapRepo: configmapRepo}
}

func (this *ConfigmapService) GetAll(ns string) *result.ErrorResult {
	return this.configmapRepo.ListAll(ns)
}
