package controller

import (
	"encoding/json"
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

	db := &db.InMemoryDB{}

	productCommandBus := bus.New()
	creatorHandler := creator.NewHandler(db)
	productCommandBus.Parse(command.ProductCommand{}, creatorHandler)

	http.HandleFunc("/product", handleProduct(productCommandBus))
}

func handleProduct(bus infraestructure.CommandBus) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlePost(bus, w, r)
			break
		default:
			w.WriteHeader(404)
		}
	}
}

func handlePost(bus infraestructure.CommandBus, w http.ResponseWriter, r *http.Request) {
	requestBody := &productRequest{}
	err := json.NewDecoder(r.Body).Decode(requestBody)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, err.Error())
		return
	}

	cmd := command.ProductCommand{
		ID:          requestBody.ID,
		Name:        requestBody.Name,
		Description: requestBody.Description,
	}
	bus.Dispatch(cmd)
	fmt.Fprintf(w, "Product created")
}

type productRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
