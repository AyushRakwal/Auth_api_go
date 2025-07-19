package app

import (
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	repo "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
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

func NewConfig() Config {
	port := config.GetString("PORT", ":8080")
	return Config{
		Address: port,
	}
}

func NewApplicaion(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

func (app *Application) Run() error {
	db, err := dbConfig.SetupDB()

	if err != nil {
		fmt.Printf("Error setting up database: %v\n", err)
		return err
	}

	ur := repo.NewUserRepository(db)
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(uc)

	server := &http.Server{
		Addr:         app.Config.Address,
		Handler:      router.SetupRouter(uRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Printf("Starting server on %s\n", app.Config.Address)
	return server.ListenAndServe()
}
