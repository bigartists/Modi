package service

import (
	"modi/core/result"
	"modi/internal/model/UserModel"
)

type IUser interface {
	GetUserList() []*UserModel.UserImpl
	GetUserDetail(id int64) *result.ErrorResult
	CreateUser(user *UserModel.UserImpl) *result.ErrorResult
	UpdateUser(id int, user *UserModel.UserImpl) *result.ErrorResult
	DeleteUser(id int) *result.ErrorResult
	SignIn(username string, password string) (*UserModel.UserImpl, error)
	SignUp(email string, username string, password string) error
}
