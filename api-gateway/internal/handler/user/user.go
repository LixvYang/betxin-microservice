package user

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lixvyang/betxin-microservice/api-gateway/internal/handler"
	"github.com/lixvyang/betxin-microservice/api-gateway/internal/service"
	"github.com/lixvyang/betxin-microservice/api-gateway/pkg/errmsg"
)

func CreateUser(ginCtx *gin.Context) {
	fmt.Println("USERADD")
	var userReq service.AddUserReq
	handler.PanicIfUserError(ginCtx.Bind(&userReq))
	// gin.Key
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	fmt.Println("userService: ", userService)
	fmt.Println(userService)
	_, err := userService.CreateUser(context.Background(), &userReq)
	handler.PanicIfUserError(err)
	handler.SendResponse(ginCtx, errmsg.SUCCSE, nil)
}
