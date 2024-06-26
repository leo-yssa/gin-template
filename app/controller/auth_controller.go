package controller

import (
	"gin-api/app/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Google(c *gin.Context)
}

type AuthControllerImpl struct {
	svc service.AuthService
}

func (u *AuthControllerImpl) Login(c *gin.Context) {
	u.svc.Login(c)
}

func (u *AuthControllerImpl) Logout(c *gin.Context) {
	u.svc.Logout(c)
}

func (u *AuthControllerImpl) Google(c *gin.Context) {
	u.svc.Google(c)
}


func NewAuthController(svc service.AuthService) AuthController {
	return &AuthControllerImpl{
		svc: svc,
	}
}