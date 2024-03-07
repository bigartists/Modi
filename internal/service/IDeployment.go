package service

import "modi/core/result"

type IDeployment interface {
	GetDeploymentsByNs(ns string) *result.ErrorResult
	IncrReplicas(ns string, dep string, dec bool) *result.ErrorResult
	GetDeploymentDetailByNsDName(ns string, dep string) *result.ErrorResult
}
