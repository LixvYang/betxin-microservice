package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lixvyang/betxin-microservice/api-gateway/pkg/errmsg"
)

// 包装错误
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		panic(err)
	}
}

func PanicIfTaskError(err error) {
	if err != nil {
		err = errors.New("taskService--" + err.Error())
		panic(err)
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: errmsg.GetErrMsg(code),
		Data:    data,
	})
}
