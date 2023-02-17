package relation

import "github.com/gin-gonic/gin"

func InitRouter(group *gin.RouterGroup) {
	group.POST("/action", rAction)
	group.GET("/follow/list", rFollowList)
	group.GET("/follower/list", rFollowerList)
	group.GET("/friend/list", rFriendList)
}
