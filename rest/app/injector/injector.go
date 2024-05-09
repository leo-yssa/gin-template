package injector

import (
	"gin-api/rest/app/config"
	"gin-api/rest/app/controller"
	"gin-api/rest/app/injector/auth"
	"gin-api/rest/app/injector/user"
	"gin-api/rest/app/middleware"
	"gin-api/rest/app/repository"

	"github.com/gin-contrib/sessions/redis"
	"gorm.io/gorm"
)

type Initialization struct {
	DB *gorm.DB
	Store redis.Store
	UserCtrl controller.UserController
	AuthCtrl controller.AuthController
}

func Init() (*Initialization, error){
	config.InitLogger()
	middleware.InitValidator()
	initialization := &Initialization{}
	if store, err := config.NewStore(); err != nil {
		return nil, err
	} else {
		initialization.Store = store
	}
	if db, err := config.NewDatabase(); err != nil {
		return nil, err
	} else {
		_ = repository.NewRoleRepository(db)
		initialization.DB = db
		initialization.UserCtrl = user.Init(db)
		initialization.AuthCtrl = auth.Init()
		return initialization, nil
	}
} 