package message

import "douyin-backend/model/user"

type Message struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	ToUserID   int64     `json:"to_user_id" gorm:"column:to_user_id"`
	FromUserID int64     `json:"from_user_id" gorm:"column:from_user_id"`
	Content    string    `json:"content" gorm:""`
	CreateTime string    `json:"create_time" gorm:""`
	ToUser     user.User `json:"-" gorm:"ForeignKey:ToUserID"`
	FromUser   user.User `json:"-" gorm:"ForeignKey:FromUserID"`
}

type Dao struct{}
