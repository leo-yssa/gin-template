package service

import (
	"context"
	"gin-api/app/config"
	"gin-api/app/constant"
	"gin-api/app/domain/dto"
	"gin-api/app/domain/error"
	"gin-api/app/middleware"
	"gin-api/app/utility"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

type AuthService interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Google(c *gin.Context)
}

type AuthServiceImpl struct {
	
}

func (u *AuthServiceImpl) Login(c *gin.Context) {
	defer error.PanicHandler(c)
	session := sessions.Default(c)
    state := middleware.GetState()
	session.Set("state", state)
	session.Options(sessions.Options{
		MaxAge: 10 * 60,
	})
    if err := session.Save(); err != nil {
		logger.Error(err)
		error.PanicException(constant.UnknownError)
	}
	config.NewOAuth().AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, config.NewOAuth().AuthCodeURL(state))
}

func (u *AuthServiceImpl) Logout(c *gin.Context) {
	session := sessions.Default(c)
 	session.Clear()
 	if err := session.Save(); err != nil {
		logger.Error(err)
		error.PanicException(constant.UnknownError)
 	}
	 c.JSON(http.StatusOK, dto.BuildResponse(constant.Success, ""))
}

func (u *AuthServiceImpl) Google(c *gin.Context) {
	defer error.PanicHandler(c)
    token, err := config.NewOAuth().Exchange(context.Background(), c.Request.FormValue("code"))
    if err != nil {
        error.PanicException(constant.InvalidRequest)
    }
	data, err := utility.Get(constant.GOOGLE_API_URL + token.AccessToken)
	if err != nil {
        error.PanicException(constant.InvalidRequest)
        return
    }
	c.JSON(http.StatusAccepted, dto.BuildResponse(constant.Success, data))
}

func NewAuthService() AuthService {
	return &AuthServiceImpl{
	}
}