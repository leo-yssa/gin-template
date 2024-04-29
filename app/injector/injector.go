package injector

import (
	"gin-api/app/config"
	"gin-api/app/controller"
	"gin-api/app/injector/auth"
	"gin-api/app/injector/user"
	"gin-api/app/repository"

	"gorm.io/gorm"
)

type Initialization struct {
	DB *gorm.DB
	UserCtrl controller.UserController
	AuthCtrl controller.AuthController
}

func Init() (*Initialization, error){
	if db, err := config.NewDatabase(); err != nil {
		return nil, err
	} else {
		_ = repository.NewRoleRepository(db)
		return &Initialization{
			DB: db,
			UserCtrl: user.Init(db),
			AuthCtrl: auth.Init(),
		}, nil
	}
} 