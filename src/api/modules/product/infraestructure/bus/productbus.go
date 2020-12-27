package bus

import (
	"reflect"

	"github.com/rlgino/go-cqrs-example/src/api/modules/shared/infraestructure"
)

type syncCommandBus struct {
	Handlers map[string]infraestructure.CommandHandler
}

// New constructor
func New() infraestructure.CommandBus {
	return &syncCommandBus{}
}

func (cmdBus *syncCommandBus) Parse(cmd infraestructure.Command, handler infraestructure.CommandHandler) {
	cmdName := reflect.TypeOf(cmd).Name()

	if cmdBus.Handlers == nil {
		cmdBus.Handlers = make(map[string]infraestructure.CommandHandler)
	}
	cmdBus.Handlers[cmdName] = handler
}

func (cmdBus *syncCommandBus) Dispatch(cmd infraestructure.Command) error {
	cmdName := reflect.TypeOf(cmd).Name()
	handler := cmdBus.Handlers[cmdName]
	handler.Invoke(cmd)
	return nil
}
