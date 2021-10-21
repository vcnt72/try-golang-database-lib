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

type PaginateProductRequest struct {
	PaginateRequest
	ProductIDs  []string `json:"product_ids"`
	ProductName string   `json:"product_name"`
	UserID      string   `json:"user_id"`
}
