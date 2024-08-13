package controllers

import (
	"github.com/gin-gonic/gin"
	"modi/config"
	"modi/src/dto"
	"modi/src/model/UserModel"
	"modi/src/result"
	"modi/src/service"
	"modi/src/utils"
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
		ret := ResultWrapper(c)(nil, err.Error())(Error)
		c.JSON(400, ret)
	}

	//// 生成 token
	prikey := []byte(config.SysYamlconfig.Jwt.PrivateKey)
	curTime := time.Now().Add(time.Second * 60 * 60 * 24)
	token, _ := utils.GenerateToken(user.Id, prikey, curTime)

	c.Set("token", token)
	ret := ResultWrapper(c)(user, "")(OK)
	c.JSON(200, ret)
}

func (a *AuthController) SignUp(c *gin.Context) {
	// 校验输入参数是否合法
	params := &dto.SignupRequest{}
	// 校验参数
	result.Result(c.ShouldBindJSON(params)).Unwrap()

	err := service.UserServiceGetter.SignUp(params.Email, params.Username, params.Password)
	if err != nil {
		ret := ResultWrapper(c)(nil, err.Error())(Error)
		c.JSON(400, ret)
	}
	ret := ResultWrapper(c)(true, "")(Created)
	c.JSON(201, ret)
}

func (a *AuthController) GetMe(c *gin.Context) {
	u := GetAuthUser(c)
	ret := ResultWrapper(c)(u, "")(OK)
	c.JSON(200, ret)
}

func GetAuthUser(c *gin.Context) *UserModel.UserImpl {
	t, exist := c.Get("auth_user")
	if !exist {
		panic("auth_user not found in gin context")
	}
	return t.(*UserModel.UserImpl)
}

func (a *AuthController) Build(r *gin.RouterGroup) {
	r.POST("/login", a.Login)
	r.POST("/register", a.SignUp)
	r.GET("/me", a.GetMe)
}

//func SetUpAuthController(r *gin.Engine) {
//	authController := NewAuthController()
//	r.POST("/login", authController.Login)
//	r.POST("/register", authController.SignUp)
//}
