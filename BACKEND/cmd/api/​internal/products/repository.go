package products

type Repository interface {
	Save(p Product) error
}

type repository struct{}

func NewRepository() Repository { return &repository{} }

func (r *repository) Save(p Product) error {
	// SQL: INSERT INTO products ...
	return nil 
}
