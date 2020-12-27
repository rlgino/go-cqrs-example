package db

import (
	"fmt"

	"github.com/rlgino/go-cqrs-example/src/api/modules/product/domain"
)

// InMemoryDB DB
type InMemoryDB struct {
	products []domain.Product
}

// Save method
func (mem *InMemoryDB) Save(prod domain.Product) error {
	if mem.products == nil {
		mem.products = []domain.Product{}
	}
	mem.products = append(mem.products, prod)

	fmt.Println("New size: ", len(mem.products))
	return nil
}

// Find method
func (mem *InMemoryDB) Find(ID string) (*domain.Product, error) {
	if mem.products == nil {
		mem.products = []domain.Product{}
	}

	return nil, nil
}
