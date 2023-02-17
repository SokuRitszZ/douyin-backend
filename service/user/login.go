package user

import (
	"douyin-backend/middleware"
	"douyin-backend/model/user"
)

type dataLogin struct {
	UserID int64
	Token  string
}

func SLogin(name string, password string) (*dataLogin, error) {
	nUser := user.User{
		Name:     name,
		Password: password,
	}
	if err := user.Dao().Login(&nUser); err != nil {
		return nil, err
	}
	token, err := middleware.ParseJWT(nUser.ID, nUser.Name)
	if err != nil {
		return nil, err
	}

	var data dataLogin
	data.UserID = nUser.ID
	data.Token = token

	return &data, nil
}
