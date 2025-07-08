package app

import (
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Address string
}

type Application struct {
	Config Config
}

func NewConfig(address string) Config {
	return Config{
		Address: address,
	}
}

func NewApplicaion(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

func (app *Application) Run() error {
	server := &http.Server{
		Addr:         app.Config.Address,
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Printf("Starting server on %s\n", app.Config.Address)
	return server.ListenAndServe()
}
