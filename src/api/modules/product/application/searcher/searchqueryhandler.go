package searcher

import (
	"github.com/rlgino/go-cqrs-example/src/api/modules/product/domain"
	"github.com/rlgino/go-cqrs-example/src/api/modules/shared/infraestructure"
)

type searcherQueryHandler struct {
	UseCase
}

// NewHandler constructor
func NewHandler(repo domain.ProductRepository) infraestructure.QueryHandler {
	return &searcherQueryHandler{
		New(repo),
	}
}

func (handler searcherQueryHandler) Invoke(queryReq infraestructure.Query) (infraestructure.Response, error) {
	prodID := domain.ProductID{
		ID: queryReq.(ProductQuery).ID,
	}
	response := handler.UseCase.Find(prodID)
	if response == nil {
		return nil, nil
	}

	ret := ProductResponse{
		ID:          response.ID.ID,
		Name:        response.Name.Name,
		Description: response.Description.Description,
	}
	return ret, nil
}
