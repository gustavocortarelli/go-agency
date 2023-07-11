package address

import (
	"github.com/gustavocortarelli/go-agency/internal/db"
	"github.com/gustavocortarelli/go-agency/internal/model"
)

func GetCountryDB() ([]model.Country, error) {
	var countries []model.Country
	err := db.R.GetSession().Preload("Cities").Find(&countries).Error

	return countries, err
}

func GetCityDB() ([]model.City, error) {
	var cities []model.City
	err := db.R.GetSession().Preload("Country").Find(&cities).Error
	return cities, err
}

type CityAndCountry struct {
	City    string `json:"city""`
	Country string `json:"country""`
}

func GetCitiesAndCountries() (any, error) {
	results, err := db.R.GetSession().
		Model(model.City{}).InnerJoins("Country").
		Select("\"Country\".name as country, city.name as city").
		Rows()
	return results, err
}
