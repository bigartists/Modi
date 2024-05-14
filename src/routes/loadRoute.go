package routes

import (
	"github.com/gin-gonic/gin"
	controllers2 "modi/src/controllers"
)

func Build(r *gin.Engine) {
	group := r.Group("/modi/v1") // *gin.RouterGroup
	controllers2.NewUserHandler().Build(group)
	controllers2.NewAuthController().Build(group)
	controllers2.NewDeploymentHandler().Build(group)
	controllers2.NewSecretController().Build(group)
}
