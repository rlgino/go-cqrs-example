package infraestructure

// CommandBus mean the bus of the command executions
type CommandBus interface {
	Parse(Command, CommandHandler)
	Dispatch(Command) error
}

// Command it's the data of the command petition
type Command interface{}

// CommandHandler parse and execute the command petition
type CommandHandler interface {
	Invoke(Command) error
}
