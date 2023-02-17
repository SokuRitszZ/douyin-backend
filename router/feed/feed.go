package feed

import (
	"douyin-backend/utils"
	"github.com/gin-gonic/gin"
)

func rFeed(c *gin.Context) {
	c.JSON(200, utils.Message{
		Code: 0,
		Msg:  "/feed",
	})
}
