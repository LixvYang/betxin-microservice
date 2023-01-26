package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(service ...any) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(service...)
		c.Keys = make(map[string]any)
		c.Keys["user"] = service[0]
		c.Next()
	}
}
