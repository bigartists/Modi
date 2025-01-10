package repo

import (
	"fmt"
	"github.com/bigartists/Modi/src/model/UserModel"
	"gorm.io/gorm"
)

type IUserRepo interface {
	FindUserAll() []*UserModel.UserImpl
	FindUserById(id int64, user *UserModel.UserImpl) (*UserModel.UserImpl, error)
	FindUserByUsername(username string) (*UserModel.UserImpl, error)
	FindUserByEmail(email string) (*UserModel.UserImpl, error)
	CreateUser(user *UserModel.UserImpl) error
	UpdateUser(id int, user *UserModel.UserImpl) error
	DeleteUser(id int) error
}

type IUserGetterImpl struct {
	db *gorm.DB
}

func NewIUserGetterImpl(db *gorm.DB) IUserRepo {
	return &IUserGetterImpl{db: db}
}

func (this *IUserGetterImpl) FindUserByUsername(username string) (*UserModel.UserImpl, error) {
	var user UserModel.UserImpl
	err := this.db.Where("username=?", username).Find(&user).Error
	return &user, err
}

func (this *IUserGetterImpl) FindUserByEmail(email string) (*UserModel.UserImpl, error) {
	var user UserModel.UserImpl
	err := this.db.Where("email=?", email).Find(&user).Error
	return &user, err
}

func (this *IUserGetterImpl) CreateUser(user *UserModel.UserImpl) error {
	return this.db.Create(user).Error
}

func (this *IUserGetterImpl) UpdateUser(id int, user *UserModel.UserImpl) error {
	//TODO implement me
	panic("implement me")
}

func (this *IUserGetterImpl) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func (this *IUserGetterImpl) FindUserById(id int64, user *UserModel.UserImpl) (*UserModel.UserImpl, error) {
	//TODO implement me
	db := this.db.Where("id=?", id).Find(user)
	if db.Error != nil || db.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found, id=%d", id)
	}
	return user, nil
}

func (this *IUserGetterImpl) FindUserAll() []*UserModel.UserImpl {
	var users []*UserModel.UserImpl
	this.db.Find(&users)
	return users
}
