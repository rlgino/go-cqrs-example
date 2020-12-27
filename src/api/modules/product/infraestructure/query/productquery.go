package query

// ProductQuery query of product
type ProductQuery struct {
	ID string
}

// ProductResponse response
type ProductResponse struct {
	ID          string
	Name        string
	Description string
}
