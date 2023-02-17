package relation

import "douyin-backend/model/user"

type FollowRelation struct {
	FollowedID int64     `json:"followed_id" gorm:"column:followed_id"`
	FanID      int64     `json:"fan_id" gorm:"column:fan_id"`
	Followed   user.User `json:"-" gorm:"ForeignKey:FollowedID"`
	Fan        user.User `json:"-" gorm:"ForeignKey:FanID"`
}

type FriendRelation struct {
	Friend1ID int64     `json:"friend_1_id" gorm:"column:friend1_id"`
	Friend2ID int64     `json:"friend_2_id" gorm:"column:friend2_id"`
	Friend1   user.User `json:"-" gorm:"ForeignKey:Friend1ID"`
	Friend2   user.User `json:"-" gorm:"ForeignKey:Friend2ID"`
}

func (FollowRelation) TableName() string {
	return "follow_relation"
}

func (FriendRelation) TableName() string {
	return "friend_relation"
}

type Dao struct{}
