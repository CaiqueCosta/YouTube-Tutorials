package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	db *sql.DB
}

func SetupStorage() (*Storage, error) {
	db, err := sql.Open("mysql", "root:toor1234@/yt_go_db")

	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}
