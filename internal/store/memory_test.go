package store

import (
	"testing"

	"github.com/mrckurz/CI-CD-MCM/internal/model"
)

func TestCreateAndGet(t *testing.T) {
	s := NewMemoryStore()

	created := s.Create(model.Product{
		Name:  "Laptop",
		Price: 999.99,
	})

	if created.ID == 0 {
		t.Error("expected assigned ID")
	}

	got, err := s.GetByID(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got.Name != "Laptop" {
		t.Errorf("expected name Laptop, got %s", got.Name)
	}

	if got.Price != 999.99 {
		t.Errorf("expected price 999.99, got %f", got.Price)
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

func TestGetByIDInvalid(t *testing.T) {
	s := NewMemoryStore()

	_, err := s.GetByID(999)

	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func TestUpdate(t *testing.T) {
	s := NewMemoryStore()

	created := s.Create(model.Product{
		Name:  "Phone",
		Price: 499.99,
	})

	updated, err := s.Update(created.ID, model.Product{
		Name:  "Updated Phone",
		Price: 599.99,
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if updated.ID != created.ID {
		t.Errorf("expected ID %d, got %d", created.ID, updated.ID)
	}

	if updated.Name != "Updated Phone" {
		t.Errorf("expected updated name, got %s", updated.Name)
	}

	if updated.Price != 599.99 {
		t.Errorf("expected updated price, got %f", updated.Price)
	}
}

func TestDeleteExisting(t *testing.T) {
	s := NewMemoryStore()

	created := s.Create(model.Product{
		Name:  "Tablet",
		Price: 299.99,
	})

	err := s.Delete(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = s.GetByID(created.ID)
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound after delete, got %v", err)
	}
}
