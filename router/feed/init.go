package feed

import "github.com/gin-gonic/gin"

func InitRouter(group *gin.RouterGroup) {
	group.GET("/", rFeed)
}
