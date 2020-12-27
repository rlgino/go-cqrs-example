package infraestructure

// QueryBus it's the query of the main query bus
type QueryBus interface {
	Parse(Query, QueryHandler)
	Invoke(Query) (Response, error)
}

// Query represents the query of the data
type Query interface{}

// QueryHandler handler the order of data
type QueryHandler interface {
	Invoke(Query) (Response, error)
}

// Response of the query petition
type Response interface{}
