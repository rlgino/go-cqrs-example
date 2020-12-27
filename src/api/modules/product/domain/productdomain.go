package domain

// Product model
type Product struct {
	ID          ProductID
	Name        ProductName
	Description ProductDescription
}

// ProductID model
type ProductID struct {
	ID string
}

// ProductName model
type ProductName struct {
	Name string
}

// ProductDescription model
type ProductDescription struct {
	Description string
}

// ProductRepository repository of product
type ProductRepository interface {
	Save(Product) error
	Find(ID ProductID) (*Product, error)
}
