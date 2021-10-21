package request

import "time"

type CreateOrderSummaryDTO struct {
	PaymentMethodID string `json:"payment_method_id"`
	UserID          string `json:"user_id"`
	OrderItems      []CreateOrderItemDTO
}

type CreateOrderItemDTO struct {
	ProductID string
	Quantity  int
}

type OrderSummaryFilterDTO struct {
	StartDate time.Time
	EndDate   time.Time
}
