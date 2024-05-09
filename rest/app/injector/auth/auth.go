//go:build wireinject
// +build wireinject

package auth

import (
	"gin-api/rest/app/controller"
	"gin-api/rest/app/service"

	"github.com/google/wire"
)

func Init() controller.AuthController{
	wire.Build(service.NewAuthService, controller.NewAuthController)
	return nil
}