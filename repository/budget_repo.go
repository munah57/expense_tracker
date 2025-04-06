package repository

import (
	"tracker/database"
	"tracker/models"
)

type BudgetRepo struct{}

type BudgetRepository interface {
	CheckDuplicateBudget(budget *models.Budget) bool
	CreateBudget(budget *models.Budget) error
	GetBudgetsByUserID(userID uint) ([]models.Budget, error)
	UpdateBudget(budget *models.Budget) error
	DeleteBudget(budget *models.Budget) error
}

func (b *BudgetRepo) CheckDuplicateBudget(budget *models.Budget) bool {
	var count int64
	database.DB.Model(&models.Budget{}).Where(&models.Budget{Category: budget.Category}).Count(&count)
	return count > 0
}
func (b *BudgetRepo) CreateBudget(budget *models.Budget) error {
	err := database.DB.Create(budget).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *BudgetRepo) GetBudgetsByUserID(UserID uint) ([]models.Budget, error) {
	var budgets []models.Budget
	err := database.DB.Where("user_id = ?", UserID).Find(&budgets).Error
	if err != nil {
		return []models.Budget{}, err
	}
	return budgets, nil
}

func (b *BudgetRepo) UpdateBudget(budget *models.Budget) error {
	err := database.DB.Save(&budget).Error
	if err !=nil {
		return err
	}
	return nil 
}

func (b *BudgetRepo) DeleteBudget(budget *models.Budget) error {
	err := database.DB.Delete(&budget).Error
	if err !=nil {
		return err
	}
	return nil 
}

