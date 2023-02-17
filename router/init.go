package router

import (
	"douyin-backend/router/comment"
	"douyin-backend/router/favorite"
	"douyin-backend/router/feed"
	"douyin-backend/router/message"
	"douyin-backend/router/publish"
	"douyin-backend/router/relation"
	"douyin-backend/router/user"
	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(group *gin.RouterGroup) {
	feed.InitRouter(group.Group("/feed"))
	user.InitRouter(group.Group("/user"))
	publish.InitRouter(group.Group("/publish"))
	favorite.InitGroup(group.Group("/favorite"))
	comment.InitRouter(group.Group("/comment"))
	relation.InitRouter(group.Group("/relation"))
	message.InitRouter(group.Group("/message"))
}
