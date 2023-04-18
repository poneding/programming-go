package repository

import (
	"errors"
	service "user/internal/service/pb"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID         uint   `gorm:"primarykey"`
	UserName       string `gorm:"unique"`
	NickName       string
	PasswordDigest string
}

const PasswordCost = 12

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}

func (u *User) CheckUserExist(req *service.UserRequest) bool {
	if err := DB.Where("user_name = ?", req.UserName).First(&u).Error; err != gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (u *User) ShowUserInfo(req *service.UserRequest) error {
	if exist := u.CheckUserExist(req); exist {
		return nil
	}
	return errors.New("user not exist")
}

func (u *User) Create(req *service.UserRequest) error {
	var user User
	var count int64

	DB.Where("user_name = ?", req.UserName).Count(&count)
	if count != 0 {
		return errors.New("user already exist")
	}
	user = User{
		UserName: req.UserName,
		NickName: req.NickName,
	}
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func BuildUser(u User) *service.UserModel {
	um := service.UserModel{
		UserID:   uint32(u.UserID),
		UserName: u.UserName,
		NickName: u.NickName,
	}
	return &um
}
