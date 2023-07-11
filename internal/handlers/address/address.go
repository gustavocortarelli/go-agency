package address

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustavocortarelli/go-agency/internal/service/address"
	"log"
	"net/http"
)

func GetCountry(context *fiber.Ctx) error {
	log.Printf("Starting get country request")
	countries, err := address.GetCountryDB()

	context.Status(http.StatusOK).JSON(countries)
	log.Printf("Get country request was done")
	return err
}

func GetCity(context *fiber.Ctx) error {
	log.Printf("Starting get city request")
	cities, err := address.GetCityDB()

	context.Status(http.StatusOK).JSON(cities)
	log.Printf("Get city request was done")
	return err
}

func GetCitiesAndCountries(context *fiber.Ctx) error {
	log.Printf("Starting get cities and country request")
	cities, err := address.GetCitiesAndCountries()

	context.Status(http.StatusOK).JSON(cities)
	log.Printf("Get cities and country request was done")
	return err
}
