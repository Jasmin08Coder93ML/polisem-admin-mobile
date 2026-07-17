package products

type Service interface {
	CreateProduct(p Product) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service { return &service{repo: r} }

func (s *service) CreateProduct(p Product) error {
	// BU ÝERDE WIZARD ALGORITMI ÝA-DA BIZNES LOGIKA BAR
	return s.repo.Save(p)
}
