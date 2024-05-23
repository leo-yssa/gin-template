package router

import (
	"gin-api/rest/app/injector"
	"gin-api/rest/app/middleware"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Init(init *injector.Initialization) *gin.Engine {
	if strings.ToLower(os.Getenv("APP_MODE")) == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(sessions.Sessions("session", init.Store))
	api := router.Group("/api")
	{
		user := api.Group("/user")
		user.POST("", init.UserCtrl.RegisterUser)
		// user.GET("", init.UserCtrl.GetAllUserData)
		// user.GET("/:userID", init.UserCtrl.GetUserById)
		// user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		// user.DELETE("/:userID", init.UserCtrl.DeleteUser)
		auth := api.Group("/auth")
		auth.GET("/login", init.AuthCtrl.Login)
		auth.Use(middleware.ValidateSession())
		{
			auth.GET("/google/callback", init.AuthCtrl.Google)
			auth.Use(middleware.ValidateJwt())
			{
				auth.POST("/logout", init.AuthCtrl.Logout)
			}
		}
		announcement := api.Group("/announcement")
		announcement.POST("", init.AnnouncementCtrl.Register)
		announcement.GET("/:announcementID", init.AnnouncementCtrl.Get)
		announcement.PUT("/:announcementID", init.AnnouncementCtrl.Modify)
		announcement.DELETE("/:announcementID", init.AnnouncementCtrl.Delete)
	}

	return router
}