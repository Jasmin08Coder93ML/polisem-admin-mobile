package products

import (
	"testing"
)

// MockRepository - bu hakyky bazanyň ýerine geçýän wagtlaýyn "mock"
type MockRepository struct {
	products map[int]Product
}

func (m *MockRepository) Save(p Product) error {
	if p.Name == "" {
		return ErrInvalidData
	}
	m.products[p.ID] = p
	return nil
}

func (m *MockRepository) GetByID(id int) (Product, error) {
	p, ok := m.products[id]
	if !ok {
		return Product{}, ErrProductNotFound
	}
	return p, nil
}

// TEST FUNKSIÝASY
func TestSaveProduct(t *testing.T) {
	// Arrange (Taýýarlyk)
	mockRepo := &MockRepository{products: make(map[int]Product)}
	newProduct := Product{ID: 1, Name: "Kabel NYM 3x2.5", Price: 150.0}

	// Act (Iş)
	err := mockRepo.Save(newProduct)

	// Assert (Barlag)
	if err != nil {
		t.Errorf("Kataloga haryt goşulanda ýalňyşlyk boldy: %v", err)
	}

	// Barlag: Haryt hakykatdan hem goşuldymy?
	saved, _ := mockRepo.GetByID(1)
	if saved.Name != "Kabel NYM 3x2.5" {
		t.Errorf("Garaşylýan: Kabel NYM 3x2.5, alnan: %s", saved.Name)
	}
}
