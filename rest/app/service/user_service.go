package service

import (
	"gin-api/rest/app/constant"
	"gin-api/rest/app/domain/dao"
	"gin-api/rest/app/domain/dto"
	"gin-api/rest/app/middleware"
	"gin-api/rest/app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

type UserService interface {
	RegisterUser(c *gin.Context)
}

type UserServiceImpl struct {
	repo repository.UserRepository
}

func (u *UserServiceImpl) RegisterUser(c *gin.Context) {
	defer middleware.PanicHandler(c)
	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error(err)
		dto.PanicException(constant.InvalidRequest)
	}
	
	data, err := u.repo.Save(&request)
	if err != nil {
		logger.Error(err)
		dto.PanicException(constant.UnknownError)
	}
	c.JSON(http.StatusAccepted, dto.BuildResponse(constant.Success, data))
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}