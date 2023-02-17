package favorite

import (
	"douyin-backend/utils"
	"github.com/gin-gonic/gin"
)

func rList(c *gin.Context) {
	c.JSON(200, utils.Message{
		Code: 0,
		Msg:  "/favorite/list",
	})
}
