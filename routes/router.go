package routes

import (
	"tracker/handler"
	"tracker/middleware"
	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handler.UserHandler, transactionHandler *handler.TransactionHandler, budgetHandler *handler.BudgetHandler ) *mux.Router {
	r := mux.NewRouter()

	//public routes

	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userHandler.LoginUser).Methods("POST")

	//initialise protected routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	//authenticated protected routes

	protected.HandleFunc("/transaction", transactionHandler.CreateTransaction).Methods("POST")
	protected.HandleFunc("/transaction", transactionHandler.GetTransactionsByUserID).Methods("GET")
	protected.HandleFunc("/transaction/balance", transactionHandler.GetTotalBalance).Methods("GET")
	protected.HandleFunc("/budget", budgetHandler.CreateBudget).Methods("POST")
	protected.HandleFunc("/budget", budgetHandler.GetBudgetsByUserID).Methods("GET")
	




//fix type
	return r
}
