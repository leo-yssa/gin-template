package service

import (
	"context"
	"gin-api/rest/app/config"
	"gin-api/rest/app/constant"
	"gin-api/rest/app/crypto"
	"gin-api/rest/app/domain/dto"
	"gin-api/rest/app/middleware"
	"gin-api/rest/app/utility"
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
	defer middleware.PanicHandler(c)
	session := sessions.Default(c)
	if session.Get("jwt.secret") == nil {
		if _, privateKey, err := crypto.GenerateKeyPair(); err != nil {
			dto.PanicException(constant.UnknownError)
		} else {
			session.Set("jwt.secret", privateKey)
		}
	}
    state := middleware.GetState()
	session.Set("state", state)
	session.Options(sessions.Options{
		MaxAge: 10 * 60,
	})
    if err := session.Save(); err != nil {
		logger.Error(err)
		dto.PanicException(constant.UnknownError)
	}
	config.NewOAuth().AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, config.NewOAuth().AuthCodeURL(state))
}

func (u *AuthServiceImpl) Logout(c *gin.Context) {
	session := sessions.Default(c)
 	session.Clear()
 	if err := session.Save(); err != nil {
		logger.Error(err)
		dto.PanicException(constant.UnknownError)
 	}
	c.JSON(http.StatusOK, dto.BuildResponse(constant.Success, ""))
}

func (u *AuthServiceImpl) Google(c *gin.Context) {
	defer middleware.PanicHandler(c)
	session := sessions.Default(c)
	secret := session.Get("jwt.secret")
    token, err := config.NewOAuth().Exchange(context.Background(), c.Request.FormValue("code"))
    if err != nil {
        dto.PanicException(constant.InvalidRequest)
    }
	data, err := utility.Get(constant.GOOGLE_API_URL + token.AccessToken)
	if err != nil {
        dto.PanicException(constant.InvalidRequest)
    }
	grpcClient, err := utility.NewClient()
	if err != nil {
		dto.PanicException(constant.GrpcServiceFailure)
	}
	grpcResult, err := grpcClient.Login("1", "test@test.com", "test", 1)
	if err != nil {
		dto.PanicException(constant.GrpcServiceFailure)
	}
	logger.Info(grpcResult)
	tokenString, err := dto.NewClaims("1", 1, secret)
	logger.Info(tokenString)
	if err != nil {
		dto.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusAccepted, dto.BuildResponse(constant.Success, data))
}

func NewAuthService() AuthService {
	return &AuthServiceImpl{
	}
}