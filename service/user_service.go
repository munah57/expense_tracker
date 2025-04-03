package service

import (
	"tracker/repository"
)
type UserService struct {
	Repo repository.UserRepository 
}