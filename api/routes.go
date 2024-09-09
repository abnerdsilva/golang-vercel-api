package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang-vercel-api/api/controller"
	"golang-vercel-api/api/repository"
	"io"
	"net/http"
)

// NotFoundHandler trata requisições para rotas não mapeadas
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Route not found"})
}

// Handler é a função principal que o Vercel chamará
func Handler(w http.ResponseWriter, r *http.Request) {
	var userRepo = repository.NewUserRepository()
	var userController = controller.NewUserRepository(userRepo)

	router := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "OK")
		if err != nil {
			return
		}
	})

	// Definindo as rotas
	router.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	// Usar o roteador para processar a solicitação
	router.ServeHTTP(w, r)
}
