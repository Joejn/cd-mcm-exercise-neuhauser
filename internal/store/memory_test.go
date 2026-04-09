package store

import (
	"errors"
	"testing"

	"github.com/mrckurz/CI-CD-MCM/internal/model"
)

type TableTestProduct struct {
	Input  model.Product
	Output model.Product
}

var defaultProducts = []TableTestProduct{
	{
		Input:  model.Product{ID: 1, Name: "Reis", Price: 10.00},
		Output: model.Product{ID: 1, Name: "Reis", Price: 10.00},
	},
	{
		Input:  model.Product{ID: 2, Name: "Linsen", Price: 12.00},
		Output: model.Product{ID: 2, Name: "Linsen", Price: 12.00},
	},
	{
		Input:  model.Product{ID: 3, Name: "Nudeln", Price: 8.00},
		Output: model.Product{ID: 3, Name: "Nudeln", Price: 8.00},
	},
}

func TestCreateAndGet(t *testing.T) {
	memoryStore := NewMemoryStore()
	for _, productToCreate := range defaultProducts {
		memoryStore.Create(productToCreate.Input)
		product, err := memoryStore.GetByID(productToCreate.Output.ID)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if product.ID != productToCreate.Output.ID {
			t.Errorf("expected product with ID %d, got %v", productToCreate.Output.ID, product)
		}
	}
}

func TestGetAllEmpty(t *testing.T) {
	s := NewMemoryStore()
	products := s.GetAll()
	if len(products) != 0 {
		t.Errorf("expected 0 products, got %d", len(products))
	}
}

func TestDeleteNonExistent(t *testing.T) {
	s := NewMemoryStore()
	err := s.Delete(999)
	if err != ErrNotFound {
		t.Error("expected ErrNotFound when deleting non-existent product")
	}
}

// TODO: Add tests for Update, Delete of existing product, and GetByID with invalid ID

func TestUpdateProduct(t *testing.T) {
	memoryStore := NewMemoryStore()
	memoryStore.Create(model.Product{
		ID:    1,
		Name:  "Reis",
		Price: 10.00,
	})

	newName := "Linsen"
	memoryStore.Update(1, model.Product{
		ID:    1,
		Name:  newName,
		Price: 10.00,
	})

	product, err := memoryStore.GetByID(1)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if product.Name != newName {
		t.Errorf("expected name %s, got %s", newName, product.Name)
	}
}

func TestDeleteProduct(t *testing.T) {
	memoryStore := NewMemoryStore()
	memoryStore.Create(model.Product{
		ID:    1,
		Name:  "Reis",
		Price: 10.00,
	})

	err := memoryStore.Delete(1)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	_, err = memoryStore.GetByID(1)
	if err == nil {
		t.Errorf("expected error when getting deleted product, got nil")
	}
}

func TestGetByIDNotFound(t *testing.T) {
	memoryStore := NewMemoryStore()
	memoryStore.Create(model.Product{
		ID:    1,
		Name:  "Reis",
		Price: 10.00,
	})

	_, err := memoryStore.GetByID(2)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected error ErrNotFound when getting not exisiting product, got %v", err)
	}
}
