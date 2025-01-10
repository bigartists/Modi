package controllers

import (
	"github.com/bigartists/Modi/src/repo"
	"github.com/bigartists/Modi/src/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConfigMapController struct {
	configmapRepo *repo.ConfigMapRepo
}

func NewConfigMapController(configmapRepo *repo.ConfigMapRepo) *ConfigMapController {
	return &ConfigMapController{configmapRepo: configmapRepo}
}

func (this *ConfigMapController) Build(r *gin.RouterGroup) {
	r.GET("/configmaps", this.ListAll)
}

func (this *ConfigMapController) ListAll(c *gin.Context) {
	namespace := &struct {
		Namespace string `form:"ns"`
	}{}
	result.Result(c.ShouldBindQuery(namespace)).Unwrap()
	ret := ResultWrapper(c)(this.configmapRepo.ListAll(namespace.Namespace).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}
