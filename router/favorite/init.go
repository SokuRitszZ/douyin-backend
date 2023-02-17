package favorite

import "github.com/gin-gonic/gin"

func InitGroup(group *gin.RouterGroup) {
	group.POST("/action", rAction)
	group.GET("/list", rList)
}
