package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type Storage struct {
	db *pgx.Conn
}

func SetupStorage() (*Storage, error) {
	connString :="postgres://marat:12345678@localhost:5432/postgres"
	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//defer db.Close(context.Background())

	return &Storage{db: db}, nil
}

func (s *Storage) GetAllCandyNames() ([]string, error) {
	var candy string
	var candies []string
	rows, err := s.db.Query(context.Background(), "SELECT name FROM candies")
	for rows.Next() {
		rows.Scan(&candy)
		candies = append(candies, candy)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return candies, nil
}

/*
*********Todo: setup candy_shop_db keyspace in cassandra*********
CREATE KEYSPACE candy_shop_db WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};

CREATE TABLE IF NOT EXISTS candies (
candy_id uuid,
category text,
name text,
price float,
PRIMARY KEY (candy_id, category)
);

INSERT INTO candies (candy_id, category, name, price) VALUES (md5(random()::text || clock_timestamp()::text)::uuid, 'Chocolate', 'KitKat', 1.99);
INSERT INTO candies (candy_id, category, name, price) VALUES (md5(random()::text || clock_timestamp()::text)::uuid, 'Sour', 'Skittles', 1.89);

CREATE USER marat WITH PASSWORD '12345678';
GRANT ALL PRIVILEGES ON DATABASE "postgres" to marat;
GRANT ALL PRIVILEGES ON TABLE candies TO marat;
***********************************************************************
*/