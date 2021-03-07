package postgresql

import (
	"CandyShop/pkg/adding"
	"context"
	"github.com/google/uuid"
	"log"
)

func (s *Storage) AddCandy(c adding.Candy) (string, error) {
	id := uuid.New().String()
	sqlString := `INSERT INTO candies (candy_id, category, name, price) VALUES ($1, $2, $3, $4)`
	if _, err := s.db.Exec(context.Background(), sqlString, id, c.Category, c.Name, c.Price);  err != nil {
		log.Println("Error while trying to save to DB:", err)
		return "", err
	}
	return id, nil
}
