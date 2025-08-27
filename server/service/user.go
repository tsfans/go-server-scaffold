package service

import (
	"context"

	"github.com/tsfans/go/framework"
	"github.com/tsfans/go/framework/utils"
	"github.com/tsfans/go/server/model/converter"
	"github.com/tsfans/go/server/model/dto"
	"github.com/tsfans/go/server/model/po"
	"github.com/tsfans/go/server/repository"
)

func GetUserByEmail(ctx context.Context, email string) (res *dto.User, err error) {
	user, err := repository.QueryData[po.User]("email = ?", email)
	if err != nil {
		return
	}
	if user == nil {
		err = framework.NewServiceError(dto.SC_USER_NOT_FOUND, "user not found with email: %s", email)
		return
	}
	res = converter.UserToDto(user)
	return
}

func PageQueryUsers(ctx context.Context, req *dto.PageRequest) (res *dto.Page, err error) {
	users, total, err := repository.PageQueryDatas[po.User](req, nil, nil)
	if err != nil {
		return
	}
	dtos := utils.MapElement(users, converter.UserToDto)
	res = dto.NewPage(req.GetPageNum(), req.GetPageSize(), total, utils.ToAnySlice(dtos))
	return
}

func CreateUser(ctx context.Context, req *dto.CreateUser) (err error) {
	user := &po.User{
		Name:     req.Name,
		Email:    req.Email,
		Age:      req.Age,
		Birthday: req.Birthday,
	}
	err = repository.CreateData(user)
	return
}

func UpdateUser(ctx context.Context, req *dto.UpdateUser) (err error) {
	user, err := repository.QueryDataById[po.User](req.Id)
	if err != nil {
		return
	}
	if user == nil {
		err = framework.NewServiceError(dto.SC_USER_NOT_FOUND, "no userfound with id: %d", req.Id)
		return
	}
	update := map[string]any{}
	if req.Name != nil {
		update["name"] = req.Name
	}
	if req.Email != nil {
		update["email"] = req.Email
	}
	if req.Age != nil {
		update["age"] = req.Age
	}
	if req.Birthday != nil {
		update["birthday"] = req.Birthday
	}
	err = repository.UpdateDataById[po.User](update, req.Id)
	return
}

func DeleteUser(ctx context.Context, id uint) (err error) {
	err = repository.DeleteById[po.User](id)
	return
}
