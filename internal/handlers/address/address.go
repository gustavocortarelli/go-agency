package address

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustavocortarelli/go-agency/internal/service/address"
	"log"
	"net/http"
)

// GetCountry Return all countries from database
//
// @Summary		Get countries
// @Description	Return all countries from database
// @Tags		address
// @Produce		json
// @Success		200	{object}	[]model.Country
// @Failure		400	{object}	model.Error
// @Failure		404	{object}	model.Error
// @Failure		500	{object}	model.Error
// @Router		/address/country [get]
func GetCountry(context *fiber.Ctx) error {
	log.Printf("Starting get country request")
	countries, err := address.GetCountryDB()

	context.Status(http.StatusOK).JSON(countries)
	log.Printf("Get country request was done")
	return err
}

// GetCity Return all cities from database
//
// @Summary		Get cities
// @Description	Return all cities from database
// @Tags		address
// @Produce		json
// @Success		200	{object}	[]model.City
// @Failure		400	{object}	model.Error
// @Failure		404	{object}	model.Error
// @Failure		500	{object}	model.Error
// @Router		/address/city [get]
func GetCity(context *fiber.Ctx) error {
	log.Printf("Starting get city request")
	cities, err := address.GetCityDB()

	context.Status(http.StatusOK).JSON(cities)
	log.Printf("Get city request was done")
	return err
}

// GetCitiesAndCountries Return city name and country without nested objects
//
// @Summary		Get city name and country
// @Description	Return city name and country without nested objects
// @Tags		address
// @Produce		json
// @Success		200	{object}	[]model.CityAndCountry
// @Failure		400	{object}	model.Error
// @Failure		404	{object}	model.Error
// @Failure		500	{object}	model.Error
// @Router		/address/city_country [get]
func GetCitiesAndCountries(context *fiber.Ctx) error {
	log.Printf("Starting get cities and country request")
	cities, err := address.GetCitiesAndCountries()

	context.Status(http.StatusOK).JSON(cities)
	log.Printf("Get cities and country request was done")
	return err
}
