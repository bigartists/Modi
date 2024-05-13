package controllers

import (
	"github.com/gin-gonic/gin"
	"modi/config"
	"modi/pkg/dto"
	"modi/pkg/result"
	"modi/pkg/service"
	"modi/pkg/utils"
	"time"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) Login(c *gin.Context) {
	// 校验输入参数是否合法
	params := &struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	// 校验参数
	result.Result(c.ShouldBindJSON(params)).Unwrap()

	user, err := service.UserServiceGetter.SignIn(params.Username, params.Password)
	if err != nil {
		ret := ResultWrapper1(c)(nil, err.Error())(Error1)
		c.JSON(400, ret)
	}

	//// 生成 token
	prikey := []byte(config.SysYamlconfig.Jwt.PrivateKey)
	curTime := time.Now().Add(time.Second * 60 * 60 * 24)
	token, _ := utils.GenerateToken(user.Id, prikey, curTime)

	c.Set("token", token)
	ret := ResultWrapper1(c)(user, "")(OK1)
	c.JSON(200, ret)
}

func (a *AuthController) SignUp(c *gin.Context) {
	// 校验输入参数是否合法
	params := &dto.SignupRequest{}
	// 校验参数
	result.Result(c.ShouldBindJSON(params)).Unwrap()

	err := service.UserServiceGetter.SignUp(params.Email, params.Username, params.Password)
	if err != nil {
		ret := ResultWrapper1(c)(nil, err.Error())(Error1)
		c.JSON(400, ret)
	}
	ret := ResultWrapper1(c)(true, "")(Created1)
	c.JSON(201, ret)
}

func (a *AuthController) Build(r *gin.RouterGroup) {
	r.POST("/login", a.Login)
	r.POST("/register", a.SignUp)
}

//func SetUpAuthController(r *gin.Engine) {
//	authController := NewAuthController()
//	r.POST("/login", authController.Login)
//	r.POST("/register", authController.SignUp)
//}
