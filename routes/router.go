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
	protected.HandleFunc("/budgets", budgetHandler.GetBudgetsByUserID).Methods("GET")
	protected.HandleFunc("/budget", budgetHandler.UpdateBudget).Methods("PUT")
	protected.HandleFunc("/budget", budgetHandler.DeleteBudget).Methods("DELETE")

	




//fix type
	return r
}
