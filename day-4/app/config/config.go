package config

import (
	"agmc/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

type ConfigDB struct {
	User string
	Pass string
	Port string
	Host string
	Name string
}

func InitDB(config ConfigDB) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.Name)

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if e != nil {
		panic(e)
	}

	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{})
}
