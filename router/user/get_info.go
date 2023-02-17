package user

import (
	"douyin-backend/service/user"
	"github.com/gin-gonic/gin"
)

type reqGetInfo struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type resGetInfo struct {
	StatusCode int64       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	User       interface{} `json:"user"`
}

func rGetInfo(c *gin.Context) {
	var req reqGetInfo
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, resGetInfo{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	data, err := user.SGetInfo(req.UserID, req.Token)
	if err != nil {
		c.JSON(200, resGetInfo{
			StatusCode: 2,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(200, resGetInfo{
		StatusCode: 0,
		StatusMsg:  "/user/get_info",
		User:       data,
	})
}
