package searcher

import "github.com/rlgino/go-cqrs-example/src/api/modules/product/domain"

// UseCase of searcher
type UseCase interface {
	Find(domain.ProductID) *domain.Product
}

type useCase struct {
	repo domain.ProductRepository
}

// New use case constructor
func New(repo domain.ProductRepository) UseCase {
	return &useCase{
		repo,
	}
}

func (useCase useCase) Find(ID domain.ProductID) *domain.Product {
	prod, _ := useCase.repo.Find(ID)
	return prod
}
