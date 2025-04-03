package repository
import (
	"tracker/models"
	"tracker/database"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
}

type UserRepo struct {}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email =?", email).First(&user).Error
	if err != nil {
	return &models.User{}, err
	}
	return &user, nil
}

func (r *UserRepo) CreateUser(user *models.User) error {
	err := database.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
