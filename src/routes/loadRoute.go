package routes

import (
	"fmt"
	controllers "github.com/bigartists/Modi/src/controllers"
	"github.com/bigartists/Modi/src/informer"
	"github.com/bigartists/Modi/src/middlewares"
	"github.com/bigartists/Modi/src/validators"
	"github.com/gin-gonic/gin"
)

func ProvideRouter(
	informerManager *informer.InformerManager,
	authMiddleware *middlewares.AuthMiddleware,
	errorMiddleware *middlewares.ErrorHandlerMiddleware,
	authController *controllers.AuthController,
	secretController *controllers.SecretController,
	configmapController *controllers.ConfigMapController,
	deployController *controllers.DeploymentController,
	userController *controllers.UserController,
) (*gin.Engine, error) {
	//r := gin.Default()
	r := gin.Default()
	// 初始化k8s client
	err := informerManager.InitInformers()
	if err != nil {
		return nil, fmt.Errorf("failed to start informer manager: %w", err)
	}
	api := r.Group("/modi/v1") // *gin.RouterGroup
	api.Use(errorMiddleware.ErrorHandler())
	api.Use(authMiddleware.JwtAuthMiddleware())

	// 加载 validator
	validators.Build()
	//controllers.NewAuthController().Build(api)
	authController.Build(api)
	secretController.Build(api)
	configmapController.Build(api)
	deployController.Build(api)
	userController.Build(api)
	return r, nil
}
