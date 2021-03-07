package reading

type Repository interface {
	GetAllCandyNames() ([]string, error)
}

type Service interface {
	GetAllCandyNames() ([]string, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) GetAllCandyNames() ([]string, error){
	cs, err := s.r.GetAllCandyNames()
	if err != nil{
		return nil, err
	}
	return cs, nil
}