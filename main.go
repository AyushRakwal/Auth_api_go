package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
)

func main() {
	config.Load()
	cfg := app.NewConfig()
	app := app.NewApplicaion(cfg)

	app.Run()
}
