package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gustavocortarelli/go-agency/configs"
	"github.com/gustavocortarelli/go-agency/internal/db"
	"github.com/gustavocortarelli/go-agency/internal/handlers/address"
	"github.com/gustavocortarelli/go-agency/internal/handlers/costumer"
	"github.com/joho/godotenv"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/costumer/:id", costumer.Get)
	api.Get("/costumer", costumer.GetAll)
	api.Post("/costumer", costumer.Create)
	api.Put("/costumer/:id", costumer.Update)
	api.Delete("/costumer/:id", costumer.Delete)

	api.Get("/country", address.GetCountry)
	api.Get("/city", address.GetCity)
	api.Get("/city_country", address.GetCitiesAndCountries)
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	err = configs.Load()
	if err != nil {
		panic(err)
	}

	err = db.OpenConnection()
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	SetupRoutes(app)

	// insert a costumer with addresses X records
	//utils.GenerateAndInsertData(100000)

	app.Listen(fmt.Sprintf(":%s", configs.GetApiPort()))
}
