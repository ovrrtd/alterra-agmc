package main

import (
	"agmc/config"
	m "agmc/middleware"
	"agmc/route"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	config.InitDB()

	e := route.New()
	m.LogMiddleware(e)

	e.Start(":8080")
}
