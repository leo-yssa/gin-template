package models

import (
	"log"

	configs "github.com/leo-yssa/gin-template/backend/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	DB *gorm.DB
}

func NewModel() *Model{
	db, err := gorm.Open(mysql.New(configs.MysqlConfig()), &gorm.Config{})
	if err != nil {
		log.Fatal("Error open database", err)
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Auth{})
	
	model := &Model{
		DB: db,
	}
	return model
}
