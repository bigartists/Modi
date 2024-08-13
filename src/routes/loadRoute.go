package routes

import (
	"github.com/gin-gonic/gin"
	controllers "modi/src/controllers"
)

func Build(r *gin.Engine) {
	group := r.Group("/modi/v1") // *gin.RouterGroup
	controllers.NewUserHandler().Build(group)
	controllers.NewAuthController().Build(group)
	controllers.NewDeploymentHandler().Build(group)
	controllers.NewSecretController().Build(group)
	controllers.NewConfigMapController().Build(group)
}
