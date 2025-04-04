package service

import (
	"tracker/models"
	"tracker/repository"
)

type TransactionService struct {
	Repo repository.TransactionRepository
}
 
//create transaction 
func (t *TransactionService) CreateTransaction(transaction *models.Transaction) error {

	err := t.Repo.CreateTransaction(transaction)
	if err != nil {
		return err
	}
	return nil 
}

func (t *TransactionService) GetTransactionsByUserID(userID uint) ([]models.Transaction, error) {
	transactions, err := t.Repo.GetTransactionsByUserID(userID)
	if err != nil {
		return []models.Transaction{}, err
	}
	return transactions, nil 
}

//Get transaction by type, if expense or 

