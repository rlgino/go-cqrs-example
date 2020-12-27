package creator

import (
	"github.com/rlgino/go-cqrs-example/src/api/modules/product/domain"
	"github.com/rlgino/go-cqrs-example/src/api/modules/shared/infraestructure"
)

// CreateProductCommandHandler it is a command handler to create product
type createProductCommandHandler struct {
	UseCase
}

// NewHandler contructor
func NewHandler(repo domain.ProductRepository) infraestructure.CommandHandler {
	return &createProductCommandHandler{
		UseCase: New(repo),
	}
}

func (handler createProductCommandHandler) Invoke(cmd infraestructure.Command) error {
	// Create VO
	// invoke usecase
	product := cmd.(ProductCommand)
	id := domain.ProductID{
		ID: product.ID,
	}
	name := domain.ProductName{
		Name: product.Name,
	}
	desc := domain.ProductDescription{
		Description: product.Description,
	}
	handler.UseCase.Persist(id, name, desc)
	return nil
}
