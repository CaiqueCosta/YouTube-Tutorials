package storage

import (
	"log"

	"github.com/gocql/gocql"
)

type Storage struct {
	db *gocql.Session
}

func SetupStorage() (*Storage, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "candy_shop_db"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		return &Storage{}, err
	}
	return &Storage{db: session}, nil
}

func (s *Storage) GetAllCandyNames() ([]string, error) {
	var candy string
	var candies []string
	iter := s.db.Query(`SELECT name FROM candies`).Iter()
	for iter.Scan(&candy) {
		candies = append(candies, candy)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
		return nil, err
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

INSERT INTO candies (candy_id, category, name, price) VALUES (now(), 'Chocolate', 'KitKat', 1.99);
INSERT INTO candies (candy_id, category, name, price) VALUES (now(), 'Sour', 'Skittles', 1.89);
*****************************************************************
*/
