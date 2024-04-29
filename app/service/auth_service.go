package service

import (
	"crypto/rand"
	"encoding/base64"
	"gin-api/app/config"
	"gin-api/app/constant"
	"gin-api/app/domain/dto"
	"gin-api/app/domain/error"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type AuthService interface {
	Login(c *gin.Context)
}

type AuthServiceImpl struct {
	
}

var store = sessions.NewCookieStore([]byte("secret"))

func (u *AuthServiceImpl) Login(c *gin.Context) {
	defer error.PanicHandler(c)
	session, _ := store.Get(c.Request, "session")
    session.Options = &sessions.Options{
        Path:   "/api/auth/login",
        MaxAge: 300,
    }
	b := make([]byte, 32)
    rand.Read(b)
    state := base64.StdEncoding.EncodeToString(b)
    session.Values["state"] = state
    session.Save(c.Request, c.Writer)
	config.NewOAuth().AuthCodeURL(state)
	c.JSON(http.StatusOK, dto.BuildResponse(constant.Success, config.NewOAuth().AuthCodeURL(state)))
}

func NewAuthService() AuthService {
	return &AuthServiceImpl{
	}
}