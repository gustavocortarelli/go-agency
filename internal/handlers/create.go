package handlers

import (
	models2 "agency/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {

	var costumer models2.Costumer

	err := json.NewDecoder(r.Body).Decode(&costumer)
	if err != nil {
		log.Printf("Error during parser: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := models2.Insert(costumer)

	var resp map[string]any

	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error during insert process: %v", err),
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf(""),
			"Id":      id,
		}
	}
	json.NewEncoder(w).Encode(resp)
}
