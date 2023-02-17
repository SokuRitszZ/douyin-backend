package user

import "github.com/gin-gonic/gin"

func InitRouter(group *gin.RouterGroup) {
	group.POST("/register", rRegister)
	group.POST("/login", rLogin)
	group.GET("/", rGetInfo)
}
