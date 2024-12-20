package handlers

import (
	"encoding/json"
	"net/http"
	"set-cover-service/algoritmo"
)

type CoverRequest struct {
	UniversalSet []int     `json:"universal_set"` // Conjunto universal
	Sets         [][]int   `json:"sets"`         // Colección de subconjuntos
}

type CoverResponse struct {
	SelectedSets [][]int `json:"selected_sets"` // Subconjuntos seleccionados
}

func CoverHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var req CoverRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	selectedSets := algoritmo.MinimumSetCover(req.UniversalSet, req.Sets)
	response := CoverResponse{SelectedSets: selectedSets}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
