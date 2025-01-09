package routes

import (
	controllers "github.com/bigartists/Modi/src/controllers"
	"github.com/gin-gonic/gin"
)

func Build(r *gin.Engine) {
	group := r.Group("/modi/v1") // *gin.RouterGroup
	controllers.NewUserHandler().Build(group)
	controllers.NewAuthController().Build(group)
	controllers.NewDeploymentHandler().Build(group)
	controllers.NewSecretController().Build(group)
	controllers.NewConfigMapController().Build(group)
}
