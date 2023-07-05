package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gustavocortarelli/go-agency/configs"
	handlers2 "github.com/gustavocortarelli/go-agency/internal/handlers"
	"net/http"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Post("/", handlers2.Create)
	router.Put("/{id}", handlers2.Update)
	router.Delete("/{id}", handlers2.Delete)
	router.Get("/", handlers2.GetAll)
	router.Get("/{id}", handlers2.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetApiPort()), router)
}
