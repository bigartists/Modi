package middlewares

import (
	"github.com/bigartists/Modi/src/controllers"
	"github.com/gin-gonic/gin"
)

type ErrorHandlerMiddleware struct {
}

func NewErrorHandlerMiddleware() *ErrorHandlerMiddleware {
	return &ErrorHandlerMiddleware{}
}

func (this *ErrorHandlerMiddleware) ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var errMsg string
				switch errTyped := err.(type) {
				case string:
					errMsg = errTyped
				case error:
					errMsg = errTyped.Error()
				default:
					errMsg = "internal Server error occurred errorHandler"
				}
				//context.JSON(500, gin.H{"error": err})
				ret := controllers.ResultWrapper(context)(nil, errMsg)(controllers.Error)
				context.JSON(500, ret)
			}
		}()
		context.Next()
	}
}
