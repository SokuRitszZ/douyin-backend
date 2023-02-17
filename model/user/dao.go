package user

import (
	"douyin-backend/model"
	"douyin-backend/utils"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	ID              int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string `json:"name" gorm:"not null"`
	Password        string `json:"-" gorm:"not null"`
	Avatar          string `json:"avatar" gorm:"default:'https://sdfsdf.dev/100x100.png'"`
	BackgroundImage string `json:"background_image" gorm:"default:'https://sdfsdf.dev/100x200.png'"`
	Signature       string `json:"signature" gorm:""`
}

type DAO struct{}

func Dao() *DAO {
	return &DAO{}
}

// Register 注册用户
func (*DAO) Register(user *User) error {
	if err := model.Db.Create(user).Error; err != nil {
		utils.Log("[User]" + err.Error())
		return err
	}
	return nil
}

// Login 登录用户
func (*DAO) Login(user *User) error {
	if err := model.Db.First(user).Error; err != nil {
		utils.Log("[User]" + err.Error())
		return err
	}
	return nil
}

func (*DAO) GetInfo(id int64) (*User, error) {
	var user User
	err := model.Db.Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		utils.Log("[User]" + err.Error())
		return nil, err
	}
	return &user, nil
}
