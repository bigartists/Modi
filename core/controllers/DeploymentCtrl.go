package controllers

import (
	"github.com/gin-gonic/gin"
	"modi/core/result"
	"modi/internal/service"
)

type DeploymentController struct {
}

func NewDeploymentHandler() *DeploymentController {
	return &DeploymentController{}
}

func (this *DeploymentController) Build(r *gin.RouterGroup) {
	r.GET("/deployments", deploymentList) // /modi/v1/deployments?ns=infra
	r.GET("/deployment", deploymentDetail)
	r.POST("/deployment/update/scale", incrReplicas)
}

func deploymentList(c *gin.Context) {
	namespace := &struct {
		Namespace string `form:"ns" binding:"required"`
	}{}
	result.Result(c.ShouldBindQuery(namespace)).Unwrap()
	ResultWrapper(c)(service.DeploymentServiceGetter.GetDeploymentsByNs(namespace.Namespace).Unwrap(), "")(OK)
}

func deploymentDetail(c *gin.Context) {
	req := &struct {
		Namespace  string `form:"ns" binding:"required"`
		Deployment string `form:"deployment" binding:"required"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	ResultWrapper(c)(service.DeploymentServiceGetter.GetDeploymentDetailByNsDName(req.Namespace, req.Deployment).Unwrap(), "")(OK)
}

func incrReplicas(c *gin.Context) {
	req := &struct {
		Namespace  string `json:"ns" binding:"required,min=1"`
		Deployment string `json:"deployment" binding:"required,min=1"`
		Dec        bool   `json:"dec"` //是否减少一个副本
	}{}
	result.Result(c.ShouldBindJSON(req)).Unwrap()
	ResultWrapper(c)(service.DeploymentServiceGetter.IncrReplicas(req.Namespace, req.Deployment, req.Dec), "")(OK)
}
