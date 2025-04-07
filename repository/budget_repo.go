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
	CheckBudgetExists(id uint) bool 
	DeleteBudget(id uint) error
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
	err := database.DB.Save(budget).Error
	if err !=nil {
		return err
	}
	return nil 
}


func (b *BudgetRepo) CheckBudgetExists(id uint) bool {
	var count int64
	database.DB.Model(&models.Budget{}).Where("id = ?", id).Count(&count)
	return count > 0
}
func (b *BudgetRepo) DeleteBudget(id uint) error {
	err := database.DB.Where("id = ?", id).Delete(&models.Budget{}).Error
	if err !=nil { //for soft delete otherwise if we use the Delete() method it will delete the record from the database entirely
		return err
	}
	return nil 
}

//after debugging, i relaised i used the wrong order to delete from database where method should be first 

/*func (b *BudgetRepo) SoftDeleteBudget(id uint) error {
	err := database.DB.Model(&models.Budget{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	if err !=nil { //for soft delete otherwise if we use the Delete() method it will delete the record from the database entirely
		return err
	}
	return nil 
}
*/
/*

When trying to perform a soft delete, not use the `Delete()` method directly, as it tries to delete the record from the database entirely rather than soft delete it.
Instead, use an `Update` to set the `deleted_at` field, while ensuring that you also provide a `WHERE` clause to filter the operation correctly. 
getting where conditions error. Added the follwing for a soft delete change to .Update("deleted_at", time.Now()).Error
*/