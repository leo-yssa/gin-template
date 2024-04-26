package injector

import (
	"gin-api/app/config"
	"gin-api/app/controller"
	"gin-api/app/injector/user"

	"gorm.io/gorm"
)

type Initialization struct {
	DB *gorm.DB
	UserCtrl controller.UserController
}

func Init() (*Initialization, error){
	if db, err := config.NewDatabase(); err != nil {
		return nil, err
	} else {
		return &Initialization{
			DB: db,
			UserCtrl: user.Init(db),
		}, nil
	}
} 