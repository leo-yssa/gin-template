package controller

import (
	"gin-api/app/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(c *gin.Context)
}

type UserControllerImpl struct {
	svc service.UserService
}

func (u *UserControllerImpl) RegisterUser(c *gin.Context) {
	u.svc.RegisterUser(c)
}

func NewUserController(svc service.UserService) UserController {
	return &UserControllerImpl{
		svc: svc,
	}
}