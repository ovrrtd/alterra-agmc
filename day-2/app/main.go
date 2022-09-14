package main

import (
	"day-two/config"
	"day-two/route"
)

func main() {
	config.InitDB()

	e := route.New()

	e.Start(":8080")
}
