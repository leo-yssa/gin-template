package injector

import (
	"gin-api/app/config"
	"gin-api/app/controller"
	"gin-api/app/injector/auth"
	"gin-api/app/injector/user"
	"gin-api/app/repository"

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