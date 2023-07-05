package models

import (
	"encoding/json"
	"github.com/gustavocortarelli/go-agency/internal/db"
	"log"
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

func Get(id int64) (costumer Costumer, err error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT id, name, surname, doc_number, birthdate FROM costumer WHERE id = $1`, id)
	err = row.Scan(&costumer.ID, &costumer.Name, &costumer.Surname, &costumer.DocNumber, &costumer.Birthdate)

	return
}

func GetAll() (costumers []Costumer, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT id, name, surname, doc_number, birthdate FROM costumer`)
	if err != nil {
		return
	}
	for rows.Next() {
		var costumer Costumer
		err = rows.Scan(&costumer.ID, &costumer.Name, &costumer.Surname, &costumer.DocNumber, &costumer.Birthdate)
		if err != nil {
			return
		}
		costumers = append(costumers, costumer)
	}

	return
}

func Insert(costumer Costumer) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	log.Printf("Data: %s", costumer)
	sql := `INSERT INTO costumer (name, surname, doc_number, birthdate) VALUES ($1, $2, $3, $4) RETURNING id`
	err = conn.QueryRow(sql, costumer.Name, costumer.Surname, costumer.DocNumber, costumer.Birthdate.GetDate()).Scan(&id)

	return
}

func Update(id int64, costumer Costumer) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE costumer SET name = $1, surname = $2, doc_number = $3, birthdate = $4 
                WHERE id = $5`, costumer.Name, costumer.Surname, costumer.DocNumber, costumer.Birthdate.GetDate(), id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func Delete(id int64) (int64, error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM costumer WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
