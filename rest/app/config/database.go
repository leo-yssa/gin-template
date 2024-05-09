package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	if strings.ToLower(os.Getenv("DATABASE")) == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
			os.Getenv("DATABASE_USER"), 
			os.Getenv("DATABASE_SECRET"), 
			os.Getenv("DATABASE_HOST"), 
			os.Getenv("DATABASE_PORT"), 
			os.Getenv("DATABASE_SCHEMA"), 
		)
		if db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: dsn, // data source name
				DefaultStringSize: 256, // default size for string fields
				DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
				DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
				DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
				SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		}), &gorm.Config{}); err != nil {
			return nil, err
		} else {
			db.Set("gorm:table_options", "ENGINE=InnoDB")
			return db, nil
		}
	}
	return nil, errors.New("Not supported database")
}