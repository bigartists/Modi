package controllers

import (
	"github.com/gin-gonic/gin"
	Model "modi/internal/model/SecretModel"
	"modi/internal/result"
	"modi/internal/service"
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
	ResultWrapper(c)(service.SecretServiceGetter.GetSecretByNs(namespace.Namespace).Unwrap(), "")(OK)
}

func postSecret(c *gin.Context) {
	postModel := &Model.PostSecretModel{}
	result.Result(c.ShouldBindJSON(postModel))
	ResultWrapper(c)(service.SecretServiceGetter.PostSecret(postModel, c).Unwrap(), "")(OK)
}
