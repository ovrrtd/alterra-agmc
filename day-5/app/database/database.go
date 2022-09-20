package database

import (
	"agmc/internal/model"
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
	// we use sync.Once for make sure we create connection only once
	once sync.Once
)

type ConfigDBType int

type ConfigDB struct {
	User string
	Pass string
	Port string
	Host string
	Name string
}

func CreateConnection() {
	config := ConfigDB{
		User: os.Getenv("APP_DB_USER"),
		Pass: os.Getenv("APP_DB_PASS"),
		Port: os.Getenv("APP_DB_PORT"),
		Host: os.Getenv("APP_DB_HOST"),
		Name: os.Getenv("APP_DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.Name)

	var e error

	once.Do(func() {
		DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	})

	if e != nil {
		panic(e)
	}

}

func GetConnection() *gorm.DB {

	if DB == nil {
		CreateConnection()
	}
	return DB
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{})
}
