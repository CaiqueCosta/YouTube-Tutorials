package database

import (
	"FirstApp/pkg/Users"
	"fmt"
)

func (s *Storage) SaveUser(u Users.User) error {
	stmt, err := s.db.Prepare("INSERT INTO users (id, first_name, last_name, email) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	r, err := stmt.Exec(nil, u.FirstName, u.LastName, u.Email)

	id, err := r.LastInsertId()
	fmt.Println(id)

	if err != nil {
		return err
	}

	return nil
}
