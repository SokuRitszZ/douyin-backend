package comment

import (
	"douyin-backend/utils"
	"github.com/gin-gonic/gin"
)

func rAction(c *gin.Context) {
	c.JSON(200, utils.Message{
		Code: 0,
		Msg:  "/comment/action",
	})
}
