package converter

import (
	"github.com/tsfans/go/server/model/dto"
	"github.com/tsfans/go/server/model/po"
)

func UserToDto(user *po.User) *dto.User {
	return &dto.User{
		Common:   *CommonToDto(&user.Model),
		Name:     user.Name,
		Email:    user.Email,
		Age:      user.Age,
		Birthday: user.Birthday,
	}
}
