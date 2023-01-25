package handler

import (
	"context"

	service "github.com/lixvyang/betxin-microservice/internal/service/pb"
)

type UserService struct{}

func NewUserSerivce() *UserService {
	return &UserService{}
}

func (*UserService) CreateUser(c context.Context, req *service.AddUserReq) (resp *service.AddUserResp, err error) {
	// var user repository.User
	// resp = new(service.AddUserResp)
	// code := repository.CheckUser(req.IdentityNumber)
	// if code != errmsg.SUCCSE {
	// 	v1.SendResponse(c, errmsg.ERROR_CATENAME_USED, nil)
	// 	return
	// }
	// copier.Copy(&user, req)

	// if code = repository.CreateUser(&user); code != errmsg.SUCCSE {
	// 	v1.SendResponse(c, errmsg.ERROR, nil)
	// 	return
	// }

	// betixnredis.DelKeys(v1.USER_LIST, v1.USER_TOTAL)
	// v1.SendResponse(c, errmsg.SUCCSE, nil)
	return
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
