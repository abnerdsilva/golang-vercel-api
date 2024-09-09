package handler

import (
	"encoding/json"
	"golang-vercel-api/api/repository"
	"io"
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

// Handler é a função principal que o Vercel chamará
func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "OK")
		if err != nil {
			return
		}
	})

	// Definindo as rotas
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUserByID).Methods("GET")

	// Usar o roteador para processar a solicitação
	router.ServeHTTP(w, r)
}
