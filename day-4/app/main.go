package main

import (
	"agmc/config"
	m "agmc/middleware"
	"agmc/route"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	cfg := config.ConfigDB{
		User: os.Getenv("APP_DB_USER"),
		Pass: os.Getenv("APP_DB_PASS"),
		Port: os.Getenv("APP_DB_PORT"),
		Host: os.Getenv("APP_DB_HOST"),
		Name: os.Getenv("APP_DB_NAME"),
	}

	config.InitDB(cfg)

	e := route.New()
	m.LogMiddleware(e)

	e.Start(":8080")
}
