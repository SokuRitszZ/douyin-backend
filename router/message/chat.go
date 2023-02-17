package message

import (
	"douyin-backend/utils"
	"github.com/gin-gonic/gin"
)

func rChat(c *gin.Context) {
	c.JSON(200, utils.Message{
		Code: 0,
		Msg:  "/message/chat",
	})
}
