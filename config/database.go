package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func ConnectDatabase() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("sqlite3", "./infra/product.sql")
		if err != nil {
			log.Fatalf("Erro ao se concetar com o Banco de dados: %v", err)
		}
	}

	return db
}
