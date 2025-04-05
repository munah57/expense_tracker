package repository

import (
	"tracker/models"
	"tracker/database"
)

type BudgetRepo struct {}

type BudgetRepository interface {
	CreateBudget(budget *models.Budget) error
	GetBudgetsByUserID(userID uint) ([]models.Budget,error)
	GetBudgetByCategory(userID uint, category string) (*models.Budget, error)	
	UpdateBudget(budget *models.Budget) error
	DeleteBudget(budget *models.Budget) error
}

func (r *BudgetRepo) CreateBudget (budget *models.Budget) error {
	err := database.DB.Create(budget).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *BudgetRepo) GetBudgetsByUserID(UserID uint) ([]models.Budget, error) {
	var budgets []models.Budget
	err := database.DB.Where("user_id = ?", UserID).Find(&budgets).Error
	if err != nil {
		return []models.Budget{}, err
	}
	return budgets, nil 
}
