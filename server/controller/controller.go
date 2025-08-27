package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsfans/go/server/model/dto"
)

func InitAllServerRoute(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.GET("", GetUserByEmail)
	user.POST("", CreateUser)
	user.PATCH("", UpdateUser)
	user.DELETE("/:id", DeleteUser)
	user.POST("/page", PageQueryUsers)
}

func JSONResponse(c *gin.Context, data any, err error) {
	var rpn dto.Response
	if err != nil {
		rpn.FailWithError(err)
		c.JSON(http.StatusOK, rpn)
		return
	}
	rpn.Success(data)
	c.JSON(http.StatusOK, rpn)
}
