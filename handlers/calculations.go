package handlers

import (
	"encoding/json"
	"net/http"
	"payroll_calculator/repository"
	"payroll_calculator/repository/calculations"
)

func ListAllCalculations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	repo, err := repository.NewRepo()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Inicializacion del repositorio de calculations
	calculationRepository := calculations.CalculationRepository{Repository: repo}

	calcs, err := calculationRepository.GettingAllCalculations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(calcs)
}
