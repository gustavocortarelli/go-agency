package costumer

import (
	"github.com/gustavocortarelli/go-agency/internal/db"
	"github.com/gustavocortarelli/go-agency/internal/model"
	"log"
)

func Get(id int64) (costumer model.Costumer, err error) {
	err = db.R.GetSession().Preload("Addresses.City.Country").Take(&costumer, id).Error
	return
}

func GetAll() (costumers []model.Costumer, err error) {
	err = db.R.GetSession().Order("id asc").Find(&costumers).Error
	return costumers, err
}

func Create(costumer model.Costumer) (model.Costumer, error) {
	log.Println(costumer.Birthdate)
	err := db.R.GetSession().Create(&costumer).Error
	return costumer, err
}

func Update(costumer model.Costumer) (model.Costumer, error) {
	log.Println(costumer.Birthdate)
	err := db.R.GetSession().Updates(&costumer).Error

	return costumer, err
}

func Delete(id int64) error {
	err := db.R.GetSession().Select("Addresses").Delete(model.Costumer{ID: id}).Error
	return err
}
