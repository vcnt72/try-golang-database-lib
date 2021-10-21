package request

type CreateProductRequest struct {
	Name     string
	Quantity int
	Price    float64
	UserID   string
}

type UpdateProductRequest struct {
	Name     string
	Quantity int
	Price    float64
}
