package adding

type Service interface {
}

type Repository interface {
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r: r}
}
