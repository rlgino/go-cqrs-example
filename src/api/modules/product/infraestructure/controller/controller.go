package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rlgino/go-cqrs-example/src/api/modules/product/application/creator"
	"github.com/rlgino/go-cqrs-example/src/api/modules/product/application/searcher"
	"github.com/rlgino/go-cqrs-example/src/api/modules/product/infraestructure/bus"
	"github.com/rlgino/go-cqrs-example/src/api/modules/product/infraestructure/command"
	"github.com/rlgino/go-cqrs-example/src/api/modules/product/infraestructure/db"
	"github.com/rlgino/go-cqrs-example/src/api/modules/product/infraestructure/query"
	"github.com/rlgino/go-cqrs-example/src/api/modules/shared/infraestructure"
)

// Run http server
func Run() {

	db := &db.InMemoryDB{}

	productCommandBus := bus.New()
	creatorHandler := creator.NewHandler(db)
	productCommandBus.Parse(command.ProductCommand{}, creatorHandler)

	productQueryBus := bus.NewQueryBus()
	searcherHandler := searcher.NewHandler(db)
	productQueryBus.Parse(query.ProductQuery{}, searcherHandler)

	http.HandleFunc("/product", handleProduct(productCommandBus, productQueryBus))
}

func handleProduct(bus infraestructure.CommandBus, queryBus infraestructure.QueryBus) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlePost(bus, w, r)
			break
		case http.MethodGet:
			handleGet(queryBus, w, r)
			break
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func handleGet(queryBus infraestructure.QueryBus, w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("product_id")

	query := query.ProductQuery{
		ID: id,
	}
	res, err := queryBus.Invoke(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	if res == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response, _ := json.Marshal(res)
	fmt.Fprintf(w, string(response))
	w.Header().Set("Content-Type", "application/json")
}

func handlePost(bus infraestructure.CommandBus, w http.ResponseWriter, r *http.Request) {
	requestBody := &productRequest{}
	err := json.NewDecoder(r.Body).Decode(requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
