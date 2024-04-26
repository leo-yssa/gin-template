package service

import (
	"gin-api/app/constant"
	"gin-api/app/domain/dao"
	"gin-api/app/domain/dto"
	"gin-api/app/domain/error"
	"gin-api/app/repository"
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
	defer error.PanicHandler(c)
	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error(err)
		error.PanicException(constant.InvalidRequest)
	}
	data, err := u.repo.Save(&request)
	if err != nil {
		logger.Error(err)
		error.PanicException(constant.UnknownError)
	}
	c.JSON(http.StatusAccepted, dto.BuildResponse(constant.Success, data))
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}