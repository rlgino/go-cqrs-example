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
	return nil
}

// Find method
func (mem *InMemoryDB) Find(id domain.ProductID) (product *domain.Product, err error) {
	fmt.Println("New size: ", len(mem.products))
	if mem.products == nil {
		mem.products = []domain.Product{}
	}

	for _, prod := range mem.products {
		fmt.Println("comparing: ", prod.ID.ID, " with ", id.ID)
		if prod.ID.ID == id.ID {
			fmt.Println("Product finded!")
			product = &domain.Product{
				ID:          prod.ID,
				Name:        prod.Name,
				Description: prod.Description,
			}
			return
		}
	}
	fmt.Println("Product not finded :(")
	return nil, nil
}
