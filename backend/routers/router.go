package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/leo-yssa/gin-template/backend/controllers"
)

type Router struct {
	ctl *controllers.Controller
}

func NewRouter() *gin.Engine{
	ctl := controllers.NewController()
	r := &Router{
		ctl: ctl,
	}
	
	return r.setupRouter()
}

func (r *Router) setupRouter() *gin.Engine {
	gin.ForceConsoleColor()
	engine := gin.Default()
	admin := engine.Group("/admin") 
	{
		admin.GET("", func(c *gin.Context) {
			c.String(200, "admin")
		})
	}
	user := engine.Group("/user")
	{
		user.POST("", r.ctl.RegisterUser)
	}
	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return engine
}