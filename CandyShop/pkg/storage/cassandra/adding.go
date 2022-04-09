package cassandra

import (
	"CandyShop/pkg/adding"
	"github.com/google/uuid"
	"log"
)

func (s *Storage) AddCandy(c adding.Candy) (string, error) {
	id := uuid.New().String()
	if err := s.db.Query(`INSERT INTO candies (candy_id, category, name, price) VALUES (?, ?, ?, ?)`,
		id, c.Category, c.Name, c.Price).Exec(); err != nil {
		log.Println("Error while trying to save to DB: ", err)
		return "", err
	}
	return id, nil
}
