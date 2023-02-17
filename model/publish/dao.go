package publish

import "douyin-backend/model/user"

type Video struct {
	ID       int64     `json:"id"`
	AuthorID int64     `json:"author_id"`
	PlayUrl  string    `json:"play_url"`
	CoverUrl string    `json:"cover_url"`
	Title    string    `json:"title"`
	Author   user.User `json:"author"`
}

type Dao struct{}
