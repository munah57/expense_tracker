package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
	"tracker/middleware"
	"tracker/models"
	"tracker/service"
)

type BudgetHandler struct {
	Service *service.BudgetService
	}


	var budgets []models.Budget
	var budget models.Budget

func (h *BudgetHandler) CreateBudget(w http.ResponseWriter, r *http.Request) {
	
	var budget models.Budget
	// Decode the request body into the budget struct
	err := json.NewDecoder(r.Body).Decode(&budget)
	if err != nil {
		http.Error(w, "invalid request boddy", http.StatusBadRequest)
	}
	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "could not get user id", http.StatusInternalServerError)
		return
	}

	budget.UserID = userID

	err = h.Service.CreateBudget(&budget) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(budget)
}

func (h *BudgetHandler) GetBudgetsByUserID(w http.ResponseWriter, r *http.Request) {
	
	userid := r.URL.Query().Get("user_id")
	if userid == "" {
		http.Error(w, "user_id query parameter is missing", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(userid)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}
	fmt.Printf("Received user_id: %s\n", userid)
    fmt.Printf("Fetching budget for user_id: %d\n", idInt)

	//call the service layer 
	budgets, err = h.Service.GetBudgetsByUserID(uint(idInt)) 
	if err != nil {
		fmt.Printf("Error fetching budgets for user_id: %d, error: %v\n", idInt, err)
        http.Error(w, "failed to fetch transactions", http.StatusInternalServerError)
        return
	}

	if len(budgets) == 0 {
        http.Error(w, "no budgets found for this user", http.StatusNotFound)
        return
    }

	w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(budgets)
    if err != nil {
        http.Error(w, "failed to encode budgets", http.StatusInternalServerError)
        return
    }

}

func (h *BudgetHandler) UpdateBudget(w http.ResponseWriter, r *http.Request) {
	
	//decode the request body into the budget struct
	err := json.NewDecoder(r.Body).Decode(&budget)
	if err != nil {
		http.Error(w, "inavalid request body", http.StatusBadRequest)
		return
	}
	
	//get the budget ID from the query parameters to identify which budget to update

	//convert the id intp int then into uint
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "id query parameter is missing", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid budget ID", http.StatusBadRequest)
		return
	}
	fmt.Printf("Received budget_id: %s\n", idStr)
	fmt.Printf("Updating budget with ID: %d\n", idInt)

	budget.ID = uint(idInt)

	id, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "could not get user id", http.StatusBadRequest)
	}

	budget.UserID = id
	
	err = h.Service.UpdateBudget(&budget)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(budget)
	
}

func (h *BudgetHandler) DeleteBudget(w http.ResponseWriter, r *http.Request) {
	
	budgetID := r.URL.Query().Get("user_id")
	if budgetID == "" {
		http.Error(w, "user_id query parameter is missing", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(budgetID)
	if err != nil {
		http.Error(w, "invalid budget ID", http.StatusBadRequest)
		return
	}
	fmt.Printf("Received budget_id: %s\n", budgetID)

	err = h.Service.DeleteBudget(uint(idInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("budget deleted successfully")
}