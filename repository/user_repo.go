package repository
import (
	"tracker/models"
)

type UserRepository interface {
	GetUserEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
}