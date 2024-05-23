//go:build wireinject
// +build wireinject

package announcement

import (
	"gin-api/rest/app/controller"
	"gin-api/rest/app/repository"
	"gin-api/rest/app/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) controller.AnnouncementController{
	wire.Build(repository.NewAnnouncementRepository, service.NewAnnouncementService, controller.NewAnnouncementController)
	return nil
}