package controllers

import (
	Model "github.com/bigartists/Modi/src/model/SecretModel"
	"github.com/bigartists/Modi/src/result"
	"github.com/bigartists/Modi/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SecretController struct {
	//secretService *service.SecretService
	secretService service.ISecret
}

func NewSecretController(secretService service.ISecret) *SecretController {
	return &SecretController{secretService: secretService}
}

func (this *SecretController) Build(r *gin.RouterGroup) {
	r.GET("/secret", this.secretList)
	r.POST("/secret", this.postSecret)
	//r.GET("/secret", secretDetail)
	//r.POST("/secret/update", updateSecret)
	//r.DELETE("/secret", deleteSecret)
}

func (this *SecretController) secretList(c *gin.Context) {
	namespace := &struct {
		Namespace string `form:"ns"`
	}{}
	result.Result(c.ShouldBindQuery(namespace)).Unwrap()
	ret := ResultWrapper(c)(this.secretService.GetSecretByNs(namespace.Namespace).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func (this *SecretController) postSecret(c *gin.Context) {
	postModel := &Model.PostSecretModel{}
	result.Result(c.ShouldBindJSON(postModel))
	ret := ResultWrapper(c)(this.secretService.PostSecret(postModel, c).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}
