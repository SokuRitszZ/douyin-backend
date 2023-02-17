package relation

import (
	"douyin-backend/utils"
	"github.com/gin-gonic/gin"
)

func rFollowList(c *gin.Context) {
	c.JSON(200, utils.Message{
		Code: 0,
		Msg:  "/relation/follow/list",
	})
}
