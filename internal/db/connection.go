package db

import (
	"database/sql"
	"fmt"
	"github.com/gustavocortarelli/go-agency/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	config := configs.GetDB()

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	conn, err := sql.Open("postgres", strConn)

	if err != nil {
		//TODO: change it to error handler
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
