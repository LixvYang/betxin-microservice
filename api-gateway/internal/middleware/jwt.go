package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lixvyang/betxin-microservice/api-gateway/pkg/errmsg"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
		}
		if code != errmsg.SUCCSE {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    errmsg.GetErrMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
