package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser() (string, error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("Fetching user in UserService")
	u.userRepository.GetByID()
	return nil
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user in UserService")
	password := "hardcoded"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return err
	}
	u.userRepository.Create(
		"user1",
		"hardcoded_pass",
		hashedPassword,
	)
	return nil
}

func (u *UserServiceImpl) LoginUser() (string, error) {
	email := "temp@temp.com"
	password := "temp_pass"

	user, err := u.userRepository.GetByEmail(email)

	if err != nil {
		fmt.Println("Error fetching user by email:", err)
		return "", err
	}

	if user == nil {
		fmt.Println("No user found with the given email")
		return "", fmt.Errorf("no user found with the given email")
	}

	isPasswordValid := utils.CheckPasswordHash(password, user.Password)

	if !isPasswordValid {
		fmt.Println("Invalid password")
		return "", fmt.Errorf("invalid password")
	}

	payload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}

	fmt.Println("JWT Token generated successfully:", tokenString)
	return tokenString, nil
}
