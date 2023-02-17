package comment

import "douyin-backend/model/user"

type Comment struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	UserID     int64     `json:"user_id" gorm:"user_id"`
	Content    string    `json:"content" gorm:"content"`
	CreateDate string    `json:"create_date" gorm:"create_date"` // 格式 mm-dd
	User       user.User `json:"-" gorm:"ForeignKey:UserID"`
}

type Dao struct{}
