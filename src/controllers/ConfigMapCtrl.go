package controllers

import (
	"github.com/bigartists/Modi/src/result"
	"github.com/bigartists/Modi/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConfigMapController struct {
	configService *service.ConfigmapService
}

func NewConfigMapController(configService *service.ConfigmapService) *ConfigMapController {
	return &ConfigMapController{configService: configService}
}

func (this *ConfigMapController) Build(r *gin.RouterGroup) {
	r.GET("/configmaps", this.ListAll)
}

func (this *ConfigMapController) ListAll(c *gin.Context) {
	namespace := &struct {
		Namespace string `form:"ns"`
	}{}
	result.Result(c.ShouldBindQuery(namespace)).Unwrap()
	ret := ResultWrapper(c)(this.configService.GetAll(namespace.Namespace).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}
