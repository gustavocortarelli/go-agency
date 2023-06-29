package models

import (
	"encoding/json"
	"strings"
	"time"
)

type JsonDate time.Time

type Costumer struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Surname   string   `json:"surname"`
	DocNumber string   `json:"doc_number"`
	Birthdate JsonDate `json:"birthdate"`
}

// Implement Marshaler and Unmarshaler interface
func (j *JsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonDate(t)
	return nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// Maybe a Format function for printing your date
func (j JsonDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (j JsonDate) GetDate() string {
	return j.Format("2006-01-02")
}
