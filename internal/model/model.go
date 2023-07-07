package model

import (
	"encoding/json"
	"strings"
	"time"
)

type Costumer struct {
	ID        int64             `gorm:"primary key;autoIncrement;column:id" json:"id"`
	Name      *string           `gorm:"column:name" json:"name"`
	Surname   *string           `gorm:"column:surname" json:"surname"`
	DocNumber *string           `gorm:"column:doc_number" json:"doc_number"`
	Birthdate date              `gorm:"column:birthdate" json:"birthdate"`
	Addresses []CostumerAddress `json:"addresses,omitempty"`
}

func (Costumer) TableName() string {
	return "costumer"
}

type CostumerAddress struct {
	ID         int64   `gorm:"primary key;autoIncrement;column:id" json:"id"`
	CityID     *int64  `json:"city_id"`
	CostumerID *int64  `json:"costumer_id"`
	Address    *string `json:"address"`
	ZipCode    *string `json:"zip_code"`
	City       City    `json:"city"`
}

func (CostumerAddress) TableName() string {
	return "costumer_address"
}

type Country struct {
	ID     int64  `gorm:"primary key" json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Cities []City `json:"cities,omitempty"`
}

func (Country) TableName() string {
	return "country"
}

type City struct {
	ID        int64   `gorm:"primary key" json:"id"`
	Name      string  `json:"name"`
	CountryID int64   `json:"country_id"`
	Country   Country `json:"country,omitempty"`
}

func (City) TableName() string {
	return "city"
}

type date time.Time

// UnmarshalJSON Implement Marshaler and Unmarshaler interface
func (j *date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = date(t)
	return nil
}

func (j date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// Format Maybe a Format function for printing your date
func (j date) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (j date) String() string {
	return j.Format("2006-01-02")
}

func (j date) GetDate() string {
	return j.Format("2006-01-02")
}
