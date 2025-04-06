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
	var budgets []models.Budget

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