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
	GetTotalIncome(userID uint) (float64, error)
	GetTotalExpense(userID uint) (float64, error) 
	GetTotalBalance(userID uint) (float64, error) 
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
func (r *TransactionRepo) GetTotalIncome(userID uint) (float64, error) {
	var totalIncome float64
	err := database.DB.Model(&models.Transaction{}).Where("user_id = ? AND type = ?", userID, "income").Select("SUM(amount)").Scan(&totalIncome).Error
	if err != nil {
		return 0, err
	}
	return totalIncome, nil
}

func (r *TransactionRepo) GetTotalExpense(userID uint) (float64, error) {
	var totalExpense float64
	err := database.DB.Model(&models.Transaction{}).Where("user_id = ? AND type = ?", userID, "expense").Select("SUM(amount)").Scan(&totalExpense).Error
	if err != nil {
		return 0, err
	}
	return totalExpense, nil
}

//Using sql to get the toal balance. We use "SUM" with a case statment "CASE WHEN - this allows us to perform the condition.
//  "THEN specifies the result  
//THEN: Specifies the result if the condition evaluates to true.
// ELSE alternative: If the condition is not met, this value is used. Often set to 0 so it doesn’t affect the sum.
//END: Concludes the CASE expression.



func (r *TransactionRepo) GetTotalBalance(userID uint) (float64, error) {
	var totalBalance float64
	err := database.DB.Model(&models.Transaction{}).Where("user_id = ?", userID).Select("SUM(CASE WHEN type = 'income' THEN amount ELSE -amount END)").Scan(&totalBalance).Error
	if err != nil {
		return 0, err
	}
	return totalBalance, nil
}