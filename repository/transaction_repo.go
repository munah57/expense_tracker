package repository

import (
	"tracker/database"
	"tracker/models"
)

/*
type Transaction struct {
	gorm.Model
	UserID   uint    `json:"user_id" gorm:"not null"` //to locate who the user is
	Type     string  `json:"type" gorm:"not null"` //type of transaction -  of two categories, either income or an expense
	Category string  `json:"category" gorm:"not null"`
	Amount   float64 `json:"amount" gorm:"not null"`
	Note 	string  `json:"note"`
}

*/ 

type TransactionRepo struct {}

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	GetTransactionsByUserID(userID uint) ([]models.Transaction, error)
}

func (r *TransactionRepo) CreateTransaction(transaction *models.Transaction) error {
	err := database.DB.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil 
}

//get Transactions by type
func (r *TransactionRepo) GetTransactionsByUserID(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := database.DB.Where("user_id", userID).Find(&transactions).Error
	if err!= nil {
		return []models.Transaction{}, err
	}
	return transactions, nil 
}