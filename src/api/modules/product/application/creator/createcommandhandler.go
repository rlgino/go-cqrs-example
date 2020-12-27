package creator

import (
	"fmt"

	"github.com/rlgino/go-cqrs-example/src/api/modules/product/domain"
	"github.com/rlgino/go-cqrs-example/src/api/modules/product/infraestructure/command"
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
	product := cmd.(command.ProductCommand)
	fmt.Println("Creating product", product.ID)
	id := domain.ProductID{}
	name := domain.ProductName{}
	desc := domain.ProductDescription{}
	handler.UseCase.Persist(id, name, desc)
	return nil
}
