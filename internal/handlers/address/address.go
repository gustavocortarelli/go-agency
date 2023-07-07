package address

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustavocortarelli/go-agency/internal/service/address"
	"net/http"
)

func GetCountry(context *fiber.Ctx) error {
	countries, err := address.GetCountryDB()

	context.Status(http.StatusOK).JSON(countries)

	return err
}

func GetCity(context *fiber.Ctx) error {
	cities, err := address.GetCityDB()

	context.Status(http.StatusOK).JSON(cities)

	return err
}
