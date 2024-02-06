package routes

import (
	"github.com/gin-gonic/gin"
	"modi/core/controllers"
)

func Build(r *gin.Engine) {
	controllers.NewUserHandler().Build(r)
	controllers.NewAuthController().Build(r)
}
