package main

import (
	"tracker/config"
	"tracker/database"
)

// the .gitignore - this is important as it shows what files it does not want to track on github
//example is .env - the .env is what holds your config variables - if it falls into hackers, your application is in dancger!
func main() {


	//config file  to load .env
	config.LoadEnv()

	//connect to DB string
	database.ConnectDB() 
	
//create register handler //login handler - user 
//create transaction handler 
}
