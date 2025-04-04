package handler

// handler layer (handles (http request& reponse) and call the service layer)
//		|
// service layer (business logic and calls the repository layer)
// 		|
// repository layer (handles direct database operations)

import (
	"encoding/json"
	"net/http"
	"tracker/models"
	"tracker/service"
)
type UserHandler struct {
	Service *service.UserService
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	//call service layer

err = h.Service.RegisterUser(&user)
if err !=nil {
	http.Error(w, "could not register user", http.StatusInternalServerError)
	return
}

//reponse
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var request models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.Service.LoginUser(request)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusInternalServerError)
		return
	}

	//response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)

}
