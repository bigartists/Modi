package middlewares

import (
	"github.com/gin-gonic/gin"
	"modi/src/controllers"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//context.JSON(500, gin.H{"error": err})
				ret := controllers.ResultWrapper(context)(nil, err.(string))(controllers.Error)
				context.JSON(500, ret)
			}
		}()
		context.Next()
	}
}
