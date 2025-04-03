package handler

// handler layer (handles (http request& reponse) and call the service layer)
//		|
// service layer (business logic and calls the repository layer)
// 		|
// repository layer (handles direct database operations)

import (
	"tracker/service"
)
type UserHandler struct {
	Service *service.UserService
}