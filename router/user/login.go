package user

import (
	"douyin-backend/service/user"
	"github.com/gin-gonic/gin"
)

type reqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type resLogin struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int64  `json:"user_id"`
	Token      string `json:"token"`
}

func rLogin(c *gin.Context) {
	var req reqLogin
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, &resLogin{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	data, err := user.SLogin(req.Username, req.Password)
	if err != nil {
		c.JSON(200, &resLogin{
			StatusCode: 2,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(200, resLogin{
		StatusCode: 0,
		Token:      data.Token,
		UserID:     data.UserID,
	})
}
