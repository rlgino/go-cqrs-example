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
	ID string
}

// ProductDescription model
type ProductDescription struct {
	ID string
}

// ProductRepository repository of product
type ProductRepository interface {
	Save(Product) error
	Find(ID string) (*Product, error)
}
