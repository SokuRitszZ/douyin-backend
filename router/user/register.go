package user

import (
	"douyin-backend/service/user"
	"github.com/gin-gonic/gin"
)

type reqRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type resRegister struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int64  `json:"user_id"`
	Token      string `json:"token"`
}

func rRegister(c *gin.Context) {
	var req reqRegister

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, resRegister{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	data, err := user.SRegister(req.Username, req.Password)
	if err != nil {
		c.JSON(200, resRegister{
			StatusCode: 2,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(200, resRegister{
		0,
		"",
		data.UserID,
		data.Token,
	})
}
