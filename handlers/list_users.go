package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"payroll_calculator/repository"
)

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting listing users")
	var repo *repository.Repository
	var err error

	repo, err = repository.NewRepo()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	users, err := repo.ListAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(users)

}
