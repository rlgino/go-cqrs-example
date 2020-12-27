package controller

import (
	"fmt"
	"net/http"

	"github.com/rlgino/go-cqrs-example/src/api/modules/product/application/creator"
	"github.com/rlgino/go-cqrs-example/src/api/modules/product/infraestructure/bus"
	"github.com/rlgino/go-cqrs-example/src/api/modules/product/infraestructure/command"
	"github.com/rlgino/go-cqrs-example/src/api/modules/product/infraestructure/db"
	"github.com/rlgino/go-cqrs-example/src/api/modules/shared/infraestructure"
)

// Run http server
func Run() {

	productCommandBus := bus.New()
	db := &db.InMemoryDB{}
	handler := creator.NewHandler(db)
	productCommandBus.Parse(command.ProductCommand{}, handler)

	http.HandleFunc("/product", createProduct(productCommandBus))
}

func createProduct(bus infraestructure.CommandBus) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cmd := command.ProductCommand{
			ID:          "asdf",
			Name:        "Product Name",
			Description: "Product description",
		}

		bus.Dispatch(cmd)
		fmt.Fprintf(w, "Product created")
	}
}
