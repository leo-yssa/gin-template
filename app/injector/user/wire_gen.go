// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package user

import (
	"gin-api/app/controller"
	"gin-api/app/repository"
	"gin-api/app/service"
	"gorm.io/gorm"
)

// Injectors from user.go:

func Init(db *gorm.DB) controller.UserController {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	return userController
}
