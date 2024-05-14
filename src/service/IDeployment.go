package service

import (
	"modi/src/result"
)

type IDeployment interface {
	GetDeploymentsByNs(ns string) *result.ErrorResult
	IncrReplicas(ns string, dep string, dec bool) *result.ErrorResult
	GetDeploymentDetailByNsDName(ns string, dep string) *result.ErrorResult
	GetPods(ns string, dep string) *result.ErrorResult
	GetPodJson(ns string, pod string) *result.ErrorResult
	DeletePod(ns string, pod string) *result.ErrorResult
	GetNs() *result.ErrorResult
}
