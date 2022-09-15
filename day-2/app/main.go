package main

import (
	"agmc/config"
	"agmc/route"
)

func main() {
	config.InitDB()

	e := route.New()

	e.Start(":8080")
}
