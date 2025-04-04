package handler

import (
	"encoding/json"
	"net/http"
	"tracker/middleware"
	"tracker/models"
	"tracker/service"
)
type TransactionHandler struct {
	Service *service.TransactionService
}


func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
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

func (h *TransactionHandler) GetTransactionsByType (w http.ResponseWriter, r *http.Request) {

	var transactions []models.Transaction

	Type := r.URL.Query().Get("type")
	if Type == "" {
		http.Error(w, "type parameter is required", http.StatusBadRequest)
		return
	}
	//call the service layer 
	var err error
	
	transactions, err = h.Service.GetTransactionsByType(Type)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//response
	json.NewEncoder(w).Encode(transactions) 
}
