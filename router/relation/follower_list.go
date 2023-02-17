package relation

import (
	"douyin-backend/utils"
	"github.com/gin-gonic/gin"
)

func rFollowerList(c *gin.Context) {
	c.JSON(200, utils.Message{
		Code: 0,
		Msg:  "/relation/follower/list",
	})
}
