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
func (t *TransactionService) GetTotalIncome(userID uint) (float64, error) {
	totalIncome, err := t.Repo.GetTotalIncome(userID)
	if err != nil {
		return 0, err
	}
	return totalIncome, nil
}
func (t *TransactionService) GetTotalExpense(userID uint) (float64, error) {
	totalExpense, err := t.Repo.GetTotalExpense(userID)
	if err != nil {
		return 0, err
	}
	return totalExpense, nil
}
//check this code out, it is a bit different from the one above
func (r *TransactionService) GetTotalBalance(userID uint) (float64, error) {
	totalBalance, err := r.Repo.GetTotalBalance(userID)
	if err != nil {
		return 0, err
	}
	return totalBalance, nil
}

