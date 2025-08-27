package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tsfans/go/framework"
	"github.com/tsfans/go/server/model/dto"
	"github.com/tsfans/go/server/service"
)

func GetUserByEmail(c *gin.Context) {
	email := c.Query("email")
	if len(email) == 0 {
		JSONResponse(c, nil, framework.NewServiceError(dto.SC_INVALID_PARAMETER, "email is empty"))
		return
	}
	res, err := service.GetUserByEmail(c.Request.Context(), email)
	JSONResponse(c, res, err)
}

func CreateUser(c *gin.Context) {
	var req dto.CreateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		JSONResponse(c, nil, framework.NewServiceError(dto.SC_INVALID_PARAMETER, err.Error()))
		return
	}
	err := service.CreateUser(c.Request.Context(), &req)
	JSONResponse(c, nil, err)
}

func UpdateUser(c *gin.Context) {
	var req dto.UpdateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		JSONResponse(c, nil, framework.NewServiceError(dto.SC_INVALID_PARAMETER, err.Error()))
		return
	}
	err := service.UpdateUser(c.Request.Context(), &req)
	JSONResponse(c, nil, err)
}

func PageQueryUsers(c *gin.Context) {
	var req dto.PageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		JSONResponse(c, nil, framework.NewServiceError(dto.SC_INVALID_PARAMETER, err.Error()))
		return
	}
	res, err := service.PageQueryUsers(c.Request.Context(), &req)
	JSONResponse(c, res, err)
}

func DeleteUser(c *gin.Context) {
	var req dto.UriId
	if err := c.ShouldBindUri(&req); err != nil {
		JSONResponse(c, nil, framework.NewServiceError(dto.SC_INVALID_PARAMETER, err.Error()))
		return
	}
	err := service.DeleteUser(c.Request.Context(), req.Id)
	JSONResponse(c, nil, err)
}
