package handler

import (
	"encoding/json"
	"net/http"
	"tracker/middleware"
	"tracker/models"
	"tracker/service"
	"strconv"
	"fmt"
)
type TransactionHandler struct {
	Service *service.TransactionService
}


func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "user_id query parameter is required", http.StatusBadRequest)
		return
	}

	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "could not get user id", http.StatusInternalServerError)
		return
	}

	transaction.UserID = userID

	err = h.Service.CreateTransaction(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)

}

func (h *TransactionHandler) GetTransactionsByUserID (w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transaction

	userid := r.URL.Query().Get("user_id")
	if userid == "" {
		http.Error(w, "user_id query parameter is missing", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(userid)
	if err != nil {
		http.Error(w, "invalid post ID", http.StatusBadRequest)
		return
	}
	fmt.Printf("Received user_id: %s\n", userid)
    fmt.Printf("Fetching transactions for user id: %d\n", idInt)

	//call the service layer 
	transactions, err = h.Service.GetTransactionsByUserID(uint(idInt))
	if err != nil {
		fmt.Printf("Error fetching transactions for user_id: %d, error: %v\n", idInt, err)
        http.Error(w, "failed to fetch transactions", http.StatusInternalServerError)
        return
	}

	if len(transactions) == 0 {
        http.Error(w, "no transactions found for this user", http.StatusNotFound)
        return
    }
	//response
	w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(transactions)
    if err != nil {
        http.Error(w, "failed to encode transactions", http.StatusInternalServerError)
        return
    }
}
