package controllers

import (
	"github.com/gin-gonic/gin"
	"modi/src/result"
	"modi/src/service"
	"net/http"
)

type ConfigMapController struct {
}

func NewConfigMapController() *ConfigMapController {
	return &ConfigMapController{}
}

func (this *ConfigMapController) Build(r *gin.RouterGroup) {
	r.GET("/configmaps", this.ListAll)
}

func (*ConfigMapController) ListAll(c *gin.Context) {
	namespace := &struct {
		Namespace string `form:"ns"`
	}{}
	result.Result(c.ShouldBindQuery(namespace)).Unwrap()
	ret := ResultWrapper(c)(service.ConfigMapInstance.ListAll(namespace.Namespace).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}
