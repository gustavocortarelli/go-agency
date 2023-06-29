package handlers

import (
	"agency/internal/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error getting ID param: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	affectedRows, err := models.Delete(int64(id))

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error during nsert process: %v", err),
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if affectedRows != 1 {
		log.Printf("Error: invalid number of deleted records: %d", affectedRows)
	}

	resp = map[string]any{
		"Error":   false,
		"Message": "",
		"Id":      id,
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
