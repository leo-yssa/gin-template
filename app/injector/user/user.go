//go:build wireinject
// +build wireinject

package user

import (
	"gin-api/app/controller"
	"gin-api/app/repository"
	"gin-api/app/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) controller.UserController{
	wire.Build(repository.NewUserRepository, service.NewUserService, controller.NewUserController)
	return nil
}