package models

import (
	"agency/internal/db"
)

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
