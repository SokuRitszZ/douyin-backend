package user

import (
	"douyin-backend/middleware"
	"douyin-backend/model/user"
)

type dataRegister struct {
	Token  string
	UserID int64
}

func SRegister(name string, password string) (*dataRegister, error) {
	nUser := user.User{
		Name:     name,
		Password: password,
	}
	if err := user.Dao().Register(&nUser); err != nil {
		return nil, err
	}
	token, err := middleware.ParseJWT(nUser.ID, nUser.Name)
	if err != nil {
		return nil, err
	}
	var data dataRegister
	data.UserID = nUser.ID
	data.Token = token
	return &data, nil
}
