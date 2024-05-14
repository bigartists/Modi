package controllers

import (
	"github.com/gin-gonic/gin"
	Model "modi/src/model/SecretModel"
	"modi/src/result"
	"modi/src/service"
	"net/http"
)

type SecretController struct {
}

func NewSecretController() *SecretController {
	return &SecretController{}
}

func (this *SecretController) Build(r *gin.RouterGroup) {
	r.GET("/secret", secretList)
	r.POST("/secret", postSecret)
	//r.GET("/secret", secretDetail)
	//r.POST("/secret/update", updateSecret)
	//r.DELETE("/secret", deleteSecret)
}

func secretList(c *gin.Context) {
	namespace := &struct {
		Namespace string `form:"ns"`
	}{}
	result.Result(c.ShouldBindQuery(namespace)).Unwrap()
	ret := ResultWrapper1(c)(service.SecretServiceGetter.GetSecretByNs(namespace.Namespace).Unwrap(), "")(OK1)
	c.JSON(http.StatusOK, ret)
}

func postSecret(c *gin.Context) {
	postModel := &Model.PostSecretModel{}
	result.Result(c.ShouldBindJSON(postModel))
	ret := ResultWrapper1(c)(service.SecretServiceGetter.PostSecret(postModel, c).Unwrap(), "")(OK1)
	c.JSON(http.StatusOK, ret)
}
