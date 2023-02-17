package middleware

import (
	"douyin-backend/config"
	"douyin-backend/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type Claims struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// ParseJWT 将带有 id 与 name 的 json 文本转化为 JWT
func ParseJWT(id int64, name string) (string, error) {
	now := time.Now()
	expireTime := now.Add(3 * time.Hour)

	claims := Claims{
		id,
		name,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Its me",
		},
	}
	jwtSecret := []byte(config.JwtSecretKey)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

// DumpJWT 解析 JWT 为带有 id 和 name 的 json
func DumpJWT(token string) (*Claims, error) {
	jwtSecret := []byte(config.JwtSecretKey)
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims := tokenClaims.Claims.(*Claims); tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// JWTAuth JWT 中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(200, utils.Message{
				Code: -1,
				Msg:  "Token 无效",
			})
		}
		claims, err := DumpJWT(token)
		if err != nil {
			c.JSON(200, utils.Message{
				Code: -1,
				Msg:  err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("id", claims.ID)
	}
}
