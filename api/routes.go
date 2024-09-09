package handler

import (
	"encoding/json"
	"golang-vercel-api/api/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var userRepo = repository.NewUserRepository()

// GetAllUsers responde com a lista de todos os usuários
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := userRepo.FindAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID responde com os dados de um usuário pelo ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := userRepo.FindByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Handler é a função principal que Vercel chama
func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()

	// Definindo as rotas
	router.HandleFunc("/api/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", GetUserByID).Methods("GET")

	// Servindo as requisições
	router.ServeHTTP(w, r)
}
