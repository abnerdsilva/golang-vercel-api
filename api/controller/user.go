package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang-vercel-api/api/repository"
	"net/http"
	"strconv"
)

type IUserController interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userRepository repository.IUserRepository
}

func NewUserRepository(userRepository repository.IUserRepository) IUserController {
	return &userController{
		userRepository: userRepository,
	}
}

// GetAllUsers responde com a lista de todos os usuários
func (u *userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.userRepository.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID responde com os dados de um usuário pelo ID
func (u *userController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := u.userRepository.FindByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
