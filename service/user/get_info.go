package user

import (
	"douyin-backend/middleware"
	"douyin-backend/model/user"
	"errors"
)

type dataGetInfo struct {
	user.User
	FollowCount    int64
	FollowerCount  int64
	IsFollow       bool
	TotalFavorited int64
	WorkCount      int64
	FavoriteCount  int64
}

func SGetInfo(id int64, token string) (*dataGetInfo, error) {
	var pack dataGetInfo

	json, err := middleware.DumpJWT(token)
	if err != nil {
		return nil, errors.New("未登录账号")
	}

	mID := json.ID
	info, err := user.Dao().GetInfo(id)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, errors.New("不存在此用户")
	}

	// TODO: 将基本信息赋值给 pack
	pack.ID = id
	pack.Name = info.Name
	pack.Avatar = info.Avatar
	pack.BackgroundImage = info.BackgroundImage
	pack.Signature = info.Signature

	// TODO: 检查 id 用户关注列表是否有关注 mID
	pack.IsFollow = mID == 1

	// TODO: 获取关注者和粉丝的数量

	// TODO: 获取作品数、点赞数、获赞数

	return &pack, nil
}
