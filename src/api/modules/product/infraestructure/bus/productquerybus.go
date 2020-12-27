package bus

import (
	"reflect"

	"github.com/rlgino/go-cqrs-example/src/api/modules/shared/infraestructure"
)

type syncQueryBus struct {
	Handlers map[string]infraestructure.QueryHandler
}

// NewQueryBus constructor
func NewQueryBus() infraestructure.QueryBus {
	return &syncQueryBus{}
}

func (cmdBus *syncQueryBus) Parse(cmd infraestructure.Query, handler infraestructure.QueryHandler) {
	cmdName := reflect.TypeOf(cmd).Name()

	if cmdBus.Handlers == nil {
		cmdBus.Handlers = make(map[string]infraestructure.QueryHandler)
	}
	cmdBus.Handlers[cmdName] = handler
}

func (cmdBus *syncQueryBus) Invoke(cmd infraestructure.Query) (infraestructure.Response, error) {
	cmdName := reflect.TypeOf(cmd).Name()
	handler := cmdBus.Handlers[cmdName]
	response, err := handler.Invoke(cmd)
	return response, err
}
