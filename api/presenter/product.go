package presenter

type ProductResponse struct {
	ID       string  `json:"product_id"`
	Name     string  `json:"product_name"`
	Price    float64 `json:"product_price"`
	Quantity int     `json:"product_quantity"`
}
