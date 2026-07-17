package warehouse

import "errors"

var (
	ErrInvalidQuantity = errors.New("количество не может быть отрицательным")
	ErrNotFound        = errors.New("товар на складе не найден")
)
