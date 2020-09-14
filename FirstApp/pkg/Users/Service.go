package Users

type Repository interface {
	SaveUser(u User) error
	ReadUser(id int) (User, error)
}

type Service interface {
	SaveUser(user User) error
	ReadUser(id int) (User, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ReadUser(id int) (User, error) {
	u := User{}
	u, err := s.r.ReadUser(id)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (s *service) SaveUser(u User) error {
	if err := s.r.SaveUser(u); err != nil {
		return err
	}
	return nil
}

// *** Cassandra DB Add User Function ***

// func (s *service) AddUser(user *User) error {
// 	query := `INSERT INTO users(first_name, last_name, email) VALUES (?, ?, ?)`
// 	if err := s.r.ExecuteQuery(query, user.FirstName, user.LastName, user.Email); err != nil {
// 		return err
// 	}
// 	return nil
// }
