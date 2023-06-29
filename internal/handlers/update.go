package handlers

import (
	models2 "agency/internal/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error getting ID param: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var costumer models2.Costumer
	err = json.NewDecoder(r.Body).Decode(&costumer)
	if err != nil {
		log.Printf("Error during parser: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	affectedRows, err := models2.Update(int64(id), costumer)
	w.Header().Add("Content-Type", "application/json")
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error during update process: %v", err),
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": "",
			"Id":      id,
		}
	}
	if affectedRows != 1 {
		log.Printf("Error: invalid number of records was updated: %d", affectedRows)
	}

	json.NewEncoder(w).Encode(resp)
}
