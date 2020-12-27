package creator

import (
	"fmt"

	"github.com/rlgino/go-cqrs-example/src/api/modules/product/infraestructure/command"
	"github.com/rlgino/go-cqrs-example/src/api/modules/shared/infraestructure"
)

// CreateProductCommandHandler it is a command handler to create product
type createProductCommandHandler struct{}

// NewHandler contructor
func NewHandler() infraestructure.CommandHandler {
	return &createProductCommandHandler{}
}

func (createProductCommandHandler) Invoke(cmd infraestructure.Command) error {
	// Create VO
	// invoke usecase
	product := cmd.(command.ProductCommand)
	fmt.Println("Creating product", product.ID)
	return nil
}
