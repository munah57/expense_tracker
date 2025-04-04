package main

import (
	"fmt"
	"net/http"
	"tracker/config"
	"tracker/database"
	"tracker/handler"
	"tracker/repository"
	"tracker/routes"
	"tracker/service"
)

//create register handler //login handler - user
//create transaction handler
// the .gitignore - this is important as it shows what files it does not want to track on github
//example is .env - the .env is what holds your config variables - if it falls into hackers, your application is in dancger!
func main() {


	//config file  to load .env
	config.LoadEnv()

	//connect to DB string
	database.ConnectDB() 

	//initialise repo
	userRepo := &repository.UserRepo{}
	transactionRepo := &repository.TransactionRepo{}

	//initialise the service
	UserService := &service.UserService{Repo: userRepo}
	transactionService := &service.TransactionService{Repo: transactionRepo}


	//initialise the handler 
	userHandler := &handler.UserHandler{Service: UserService}
	transactionHandler := &handler.TransactionHandler{Service: transactionService}
	
	//define routes
	router := routes.SetupRouter(userHandler, transactionHandler)

	//start the server
	fmt.Println("server is running on localhost:8080...")
	http.ListenAndServe(":8080", router)



}
