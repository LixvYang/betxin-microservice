package handler

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/lixvyang/betxin-microservice/user/internal/repository"
	service "github.com/lixvyang/betxin-microservice/user/internal/service"
	"github.com/lixvyang/betxin-microservice/user/pkg/errmsg"
)

type UserService struct {
	*service.UnimplementedUserServiceServer
}

func NewUserSerivce() *UserService {
	return &UserService{}
}

func (*UserService) CreateUser(c context.Context, req *service.AddUserReq) (resp *service.AddUserResp, err error) {
	var user repository.User
	resp = new(service.AddUserResp)
	code := repository.CheckUser(req.IdentityNumber)
	if code != errmsg.SUCCSE {
		return resp, errors.New(errmsg.GetErrMsg(code))
	}
	copier.Copy(&user, req)

	if code = repository.CreateUser(&user); code != errmsg.SUCCSE {
		return resp, errors.New(errmsg.GetErrMsg(code))
	}
	resp.Code = errmsg.SUCCSE
	resp.Message = errmsg.GetErrMsg(code)
	return resp, nil
}

func (*UserService) UpdateUser(c context.Context, req *service.UpdateUserReq) (resp *service.UpdateUserResp, err error) {
	return
}

func (*UserService) DeleteUser(c context.Context, req *service.DelUserReq) (resp *service.DelUserResp, err error) {
	return
}

func (*UserService) GetUserInfoByUserId(c context.Context, req *service.GetUserByIdReq) (resp *service.GetUserByIdResp, err error) {
	return
}

func (*UserService) ListUser(c context.Context, req *service.ListUserReq) (resp *service.ListUserResp, err error) {
	return
}
