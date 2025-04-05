package routes

import (
	"tracker/handler"
	"tracker/middleware"
	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handler.UserHandler, transactionHandler *handler.TransactionHandler) *mux.Router {
	r := mux.NewRouter()

	//public routes

	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userHandler.LoginUser).Methods("POST")

	//initialise protected routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	//authenticated protected routes

	protected.HandleFunc("/transaction", transactionHandler.CreateTransaction).Methods("POST")
	protected.HandleFunc("/transation", transactionHandler.GetTransactionsByType).Methods("GET")
//fix type
	return r
}
