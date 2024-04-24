package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
)

type Database struct {
	Mysql *Mysql `mapstructure:"mysql"`
}

type Mysql struct {
	User string `mapstructure:"user"`
	Secret string `mapstructure:"secret"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Schema string `mapstructure:"schema"`
}

func MysqlConfig() mysql.Config {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
		Config.Database.Mysql.User, 
		Config.Database.Mysql.Secret, 
		Config.Database.Mysql.Host,
		Config.Database.Mysql.Port,
		Config.Database.Mysql.Schema,
	)
	return mysql.Config{
		DSN: dsn, // data source name
 		DefaultStringSize: 256, // default size for string fields
  		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
  		DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
  		DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
  		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}
}