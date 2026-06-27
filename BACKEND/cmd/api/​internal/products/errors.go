package products

import "errors"

// Sentinel errors - bu ýalňyşlyklary proýektiň islendik ýerinde barlap bolar.
var (
	ErrProductNotFound = errors.New("product: the requested product does not exist")
	ErrInvalidData     = errors.New("product: provided product data is invalid")
	ErrDatabase        = errors.New("product: database connection or query failure")
	ErrUnauthorized    = errors.New("product: unauthorized access to product module")
)

// Has çylşyrymly ýalňyşlyklar üçin struct ulanmak mümkin:
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return "validation error on " + e.Field + ": " + e.Message
}
