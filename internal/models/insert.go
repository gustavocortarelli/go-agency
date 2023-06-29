package models

import (
	"agency/internal/db"
	"log"
)

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
