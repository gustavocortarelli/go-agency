package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gustavocortarelli/go-agency/internal/models"
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

func Create(w http.ResponseWriter, r *http.Request) {

	var costumer models.Costumer

	err := json.NewDecoder(r.Body).Decode(&costumer)
	if err != nil {
		log.Printf("Error during parser: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := models.Insert(costumer)

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

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error getting ID param: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var costumer models.Costumer
	err = json.NewDecoder(r.Body).Decode(&costumer)
	if err != nil {
		log.Printf("Error during parser: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	affectedRows, err := models.Update(int64(id), costumer)
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
