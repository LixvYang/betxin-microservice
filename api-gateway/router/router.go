package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lixvyang/betxin-microservice/api-gateway/internal/handler/user"
	"github.com/lixvyang/betxin-microservice/api-gateway/internal/middleware"
)

func NewRouter(service ...any) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service))

	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})
		v1.POST("user/add", user.CreateUser)
	}

	return ginRouter
}
