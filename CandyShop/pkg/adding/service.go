package adding

type Repository interface {
	AddCandy(Candy) (string, error)
}

type Service interface {
	AddCandy(Candy) (string, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) AddCandy(c Candy) (string, error) {
	id, err := s.r.AddCandy(c)
	if err != nil {
		return "", err
	}
	return id, nil
}
