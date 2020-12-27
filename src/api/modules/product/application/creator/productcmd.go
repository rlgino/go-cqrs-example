package creator

// ProductCommand command to serve to product command bus
type ProductCommand struct {
	ID          string
	Name        string
	Description string
}
