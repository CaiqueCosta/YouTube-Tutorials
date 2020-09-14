package database

import "FirstApp/pkg/Users"

func (s *Storage) ReadUser(id int) (Users.User, error) {
	u := Users.User{}
	r, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		return u, err
	}

	for r.Next() {
		err = r.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email)
		if err != nil {
			panic(err)
		}
	}
	return u, nil
}
