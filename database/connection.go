package db

import (
	"api-golang/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()
	urlConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s db_name=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	connectionDb, err := sql.Open("postgres", urlConnection)
	if err != nil {
		panic(err)
	}

	err = connectionDb.Ping()
	return connectionDb, err
}
