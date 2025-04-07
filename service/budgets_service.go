package service

import (
	"tracker/models"
	"tracker/repository"
	"errors"
)

type BudgetService struct {
	Repo repository.BudgetRepository
}


//create a new budget 

func (b *BudgetService) CreateBudget(budget *models.Budget) error {
	/*ok := b.Repo.CheckDuplicateBudget(budget)
	if !ok {
		return errors.New("budget already exists")
	}

	*/
	err := b.Repo.CreateBudget(budget) 
	if err != nil {
		return err
	}
	return nil 
}

//get all budgets by user id 

func (b *BudgetService) GetBudgetsByUserID(userID uint) ([]models.Budget, error) {
	//CHECK IF USER exists 
	//
	budgets, err := b.Repo.GetBudgetsByUserID(userID)
	if err != nil {
		return []models.Budget{}, err
	}
	return budgets, nil 
}

func (b *BudgetService) UpdateBudget(budget *models.Budget) error {
	err := b.Repo.UpdateBudget(budget)
	if err != nil {
		return err
	}
	return nil
}

//delete a budget by id 
func (b *BudgetService) DeleteBudget(id uint) error {	

	ok := b.Repo.CheckBudgetExists(id)
	if !ok {
		return errors.New("budget not found")
	}
	err := b.Repo.DeleteBudget(id)
	if err != nil {
		return err
	}
	return nil
}
