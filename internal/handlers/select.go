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

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error getting ID param: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	costumer, err := models.Get(int64(id))

	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		log.Printf("Error get all costumers: %d", err)
		resp := map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error get all costumers: %v", err),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}

	json.NewEncoder(w).Encode(costumer)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	costumers, err := models.GetAll()
	if err != nil {
		log.Printf("Error during get all costumers: %d", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(costumers)
}
