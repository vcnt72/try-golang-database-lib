package order_summary

type CreateOrderSummaryDTO struct {
	PaymentMethodID string
	UserID          string
	OrderItems      []CreateOrderItemDTO
}

type CreateOrderItemDTO struct {
	ProductID string
	Quantity  int
}
