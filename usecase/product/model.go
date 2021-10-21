package product

type CreateProductDTO struct {
	Name     string
	Price    float64
	Quantity int
	UserID   string
}

type UpdateProductDTO struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}
