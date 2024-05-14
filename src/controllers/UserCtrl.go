package controllers

import (
	"github.com/gin-gonic/gin"
	"modi/src/result"
	"modi/src/service"
)

type UserController struct {
}

func NewUserHandler() *UserController {
	return &UserController{}
}

// GET /users/123

// Build Build方法
func (this *UserController) Build(r *gin.RouterGroup) {
	r.GET("/users", UserList)
	r.GET("/user/:id", UserDetail)
}

func UserList(c *gin.Context) {
	ret := ResultWrapper1(c)(service.UserServiceGetter.GetUserList(), "")(OK1)
	c.JSON(200, ret)
}

func UserDetail(c *gin.Context) {
	id := &struct {
		Id int64 `uri:"id" binding:"required"`
	}{}
	result.Result(c.ShouldBindUri(id)).Unwrap()
	ret := ResultWrapper1(c)(service.UserServiceGetter.GetUserDetail(id.Id).Unwrap(), "")(OK1)
	c.JSON(200, ret)
}

//
//func UserSave(c *gin.Context) {
//	u := UserModel.New()
//	result.Result(c.ShouldBindJSON(u)).Unwrap()
//	ResultWrapper(c)("save user", "10086", "true")(OK)
//}
