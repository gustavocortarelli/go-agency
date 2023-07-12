package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gustavocortarelli/go-agency/configs"
	_ "github.com/gustavocortarelli/go-agency/docs"
	"github.com/gustavocortarelli/go-agency/internal/db"
	"github.com/gustavocortarelli/go-agency/internal/handlers/address"
	"github.com/gustavocortarelli/go-agency/internal/handlers/costumer"
	"github.com/joho/godotenv"
	"github.com/swaggo/fiber-swagger"
	"log"
)

//	@title			GO-Agency Application
//	@version		0.2.1
//	@description	This is a GO project built to learn the GO language, and test libraries like `GORM` and `SWAGGO`
//	@contact.name	gustavocortarelli
//	@contact.url	https://github.com/gustavocortarelli/
//	@license.name	MIT License
//	@license.url	https://spdx.org/licenses/MIT.html
//	@servers.url	/api/v1/

func SetupRoutes(app *fiber.App) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	r := app.Group("/api/v1/", func(ctx *fiber.Ctx) error {
		ctx.Set("Version", "v1")
		return ctx.Next()
	})
	r.Group("/costumer").
		Get("/:id", costumer.Get).
		Get("/", costumer.GetAll).
		Post("/", costumer.Create).
		Put("/:id", costumer.Update).
		Delete("/:id", costumer.Delete)

	r.Group("/address").
		Get("/country", address.GetCountry).
		Get("/city", address.GetCity).
		Get("/city_country", address.GetCitiesAndCountries)

}

func main() {

	err := godotenv.Load(".env")

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

	err = app.Listen(fmt.Sprintf(":%s", configs.GetApiPort()))
	if err != nil {
		log.Fatalf("fiber.Listen failed %s", err)
	}
}
