package service

import (
	"github.com/gin-gonic/gin"
	"modi/src/result"
)

type IDeployment interface {
	GetDeploymentsByNs(ns string) *result.ErrorResult
	IncrReplicas(ns string, dep string, dec bool) *result.ErrorResult
	GetDeploymentDetailByNsDName(ns string, dep string) *result.ErrorResult
	GetPods(ns string, dep string) *result.ErrorResult
	GetPodJson(ns string, pod string) *result.ErrorResult
	GetPodDetail(ns string, pod string) *result.ErrorResult
	DeletePod(ns string, pod string) *result.ErrorResult
	GetNs() *result.ErrorResult
	GetPodLogs(c *gin.Context, ns string, pod string, cname string) *result.ErrorResult
	GetPodContainer(ns string, podName string) *result.ErrorResult
}
