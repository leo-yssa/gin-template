package controller

import (
	"gin-api/rest/app/service"

	"github.com/gin-gonic/gin"
)

type AnnouncementController interface {
	Delete(c *gin.Context)
	Get(c *gin.Context)
	Modify(c *gin.Context)
	Register(c *gin.Context)
}

type AnnouncementControllerImpl struct {
	svc service.AnnouncementService
}

func (a *AnnouncementControllerImpl) Delete(c *gin.Context) {
	a.svc.Delete(c)
}

func (a *AnnouncementControllerImpl) Get(c *gin.Context) {
	a.svc.Get(c)
}

func (a *AnnouncementControllerImpl) Modify(c *gin.Context) {
	a.svc.Modify(c)
}

func (a *AnnouncementControllerImpl) Register(c *gin.Context) {
	a.svc.Register(c)
}

func NewAnnouncementController(svc service.AnnouncementService) AnnouncementController {
	return &AnnouncementControllerImpl{
		svc: svc,
	}
}

