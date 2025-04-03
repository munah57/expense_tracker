package service

import (
	"errors"
	"tracker/middleware"
	"tracker/models"
	"tracker/repository"
	"tracker/utils"
)

//calls the repo service
type UserService struct {
	Repo repository.UserRepository 
}


func (s *UserService) RegisterUser(user *models.User) error {
	//check if user exist already we call on the get user by email func on 
	_, err := s.Repo.GetUserByEmail(user.Email)
	if err == nil {
		return errors.New("email already in use")
	}
	//hash the password 
	hashpass, err := utils.HashPassword(user.Password)
	if err != nil {
		return errors.New("email already in use")
	}

	user.Password = hashpass

	// call the create method
	err = s.Repo.CreateUser(user)
	if err !=nil {
		return err
	}

	return nil 
}

//check if user is in db 
//compare password 
//generate token


func (s *UserService) LoginUser(request models.LoginRequest) (string, error) {
	user, err := s.Repo.GetUserByEmail(request.Email)
	if err !=nil {
		return "", err
	}

	err = utils.ComparePassword(user.Password, request.Password)
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil 
}

