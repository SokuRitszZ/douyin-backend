package main

import (
	"douyin-backend/model"
	"douyin-backend/model/comment"
	"douyin-backend/model/favorite"
	"douyin-backend/model/message"
	"douyin-backend/model/publish"
	"douyin-backend/model/relation"
	"douyin-backend/model/user"
	"douyin-backend/router"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	engine := gin.Default()

	if err := Init(engine); err != nil {
		os.Exit(-1)
	}

	// 测试连通
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "PONG",
		})
	})

	engine.Run()
}

func Init(engine *gin.Engine) error {
	// 初始化路由
	group := engine.Group("/douyin")
	router.Init(group)

	// 初始化数据库
	if err := model.Init(); err != nil {
		return err
	}
	//InitModel()

	return nil
}

func InitModel() {
	model.Migrate(user.User{})
	model.Migrate(relation.FollowRelation{})
	model.Migrate(relation.FriendRelation{})
	model.Migrate(publish.Video{})
	model.Migrate(message.Message{})
	model.Migrate(favorite.Favorite{})
	model.Migrate(comment.Comment{})
}
