package presenter

type OrderSummaryResponse struct {
	PaymentMethod PaymentMethodResponse `json:"order_summary_payment_method"`
	User          UserResponse          `json:"order_summary_user"`
	Total         float64               `json:"order_summary_total"`
	ID            string                `json:"order_summary_id"`
	OrderItems    []OrderItemResponse   `json:"order_summary_order_items"`
}

type OrderItemResponse struct {
	ID       string          `json:"order_item_id"`
	Product  ProductResponse `json:"order_item_product"`
	Quantity int             `json:"order_item_qty"`
	Total    float64         `json:"order_item_total"`
}
