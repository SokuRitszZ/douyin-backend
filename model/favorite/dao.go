package favorite

import (
	"douyin-backend/model/publish"
	"douyin-backend/model/user"
)

type Favorite struct {
	UserID  int64         `json:"user_id" gorm:"column:user_id"`
	VideoID int64         `json:"video_id" gorm:"column:publish_id"`
	User    user.User     `json:"user" gorm:"ForeignKey:UserID"`
	Video   publish.Video `json:"video" gorm:"ForeignKey:VideoID"`
}

type Dao struct{}
