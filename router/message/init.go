package message

import "github.com/gin-gonic/gin"

func InitRouter(group *gin.RouterGroup) {
	group.POST("/action", rAction)
	group.GET("/chat", rChat)
}
