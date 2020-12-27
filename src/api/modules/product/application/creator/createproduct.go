package creator

import "github.com/rlgino/go-cqrs-example/src/api/modules/product/domain"

// UseCase of product creator
type UseCase interface {
	Persist(domain.ProductID, domain.ProductName, domain.ProductDescription)
}

type useCase struct {
	repo domain.ProductRepository
}

// New constructor
func New(repo domain.ProductRepository) UseCase {
	return &useCase{
		repo: repo,
	}
}

func (useCase useCase) Persist(ID domain.ProductID, name domain.ProductName, desc domain.ProductDescription) {
	prod := domain.Product{
		ID:          ID,
		Name:        name,
		Description: desc,
	}
	useCase.repo.Save(prod)
}
