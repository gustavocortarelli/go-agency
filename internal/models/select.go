package models

import (
	"agency/internal/db"
)

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
