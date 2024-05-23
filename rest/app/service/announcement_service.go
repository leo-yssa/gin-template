package service

import (
	"gin-api/rest/app/constant"
	"gin-api/rest/app/domain/dto"
	"gin-api/rest/app/middleware"
	"gin-api/rest/app/repository"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

type AnnouncementService interface {
	Delete(c *gin.Context)
	Get(c *gin.Context)
	Modify(c *gin.Context)
	Register(c *gin.Context)
}

type AnnouncementServiceImpl struct {
	repo repository.AnnouncementRepository
}

func (a *AnnouncementServiceImpl) Delete(c *gin.Context) {

}

func (a *AnnouncementServiceImpl) Get(c *gin.Context) {

}

func (a *AnnouncementServiceImpl) Modify(c *gin.Context) {

}

func (a *AnnouncementServiceImpl) Register(c *gin.Context) {
	defer middleware.PanicHandler(c)
	form, err := c.MultipartForm()
	if err != nil {
		dto.PanicException(constant.UnknownError)
	}
	
	files := form.File["file"]
	name := form.Value["title"]
	description := form.Value["description"]
	logger.Info(form, files, name, description)
}

func NewAnnouncementService(repo repository.AnnouncementRepository) AnnouncementService {
	return &AnnouncementServiceImpl{
		repo: repo,
	}
}