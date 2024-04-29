package controller

import (
	"gin-api/app/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
}

type AuthControllerImpl struct {
	svc service.AuthService
}

func (u *AuthControllerImpl) Login(c *gin.Context) {
	u.svc.Login(c)
}

func NewAuthController(svc service.AuthService) AuthController {
	return &AuthControllerImpl{
		svc: svc,
	}
}