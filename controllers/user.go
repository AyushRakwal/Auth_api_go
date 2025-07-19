package controllers

import (
	"AuthInGo/services"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("GetUserById called"))
	uc.UserService.GetUserById()
	w.Write([]byte("User has been fetched successfully"))
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("CreateUser called"))
	uc.UserService.CreateUser()
	w.Write([]byte("User has been created successfully"))
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("LoginUser called"))
	uc.UserService.LoginUser()
	w.Write([]byte("User has been logged in successfully"))
}
