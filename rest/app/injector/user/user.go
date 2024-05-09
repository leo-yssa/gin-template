//go:build wireinject
// +build wireinject

package user

import (
	"gin-api/rest/app/controller"
	"gin-api/rest/app/repository"
	"gin-api/rest/app/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) controller.UserController{
	wire.Build(repository.NewUserRepository, service.NewUserService, controller.NewUserController)
	return nil
}