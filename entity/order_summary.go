package entity

import "time"

type OrderSummary struct {
	ID              string
	Total           float64
	PaymentMethodID string
	PaymentMethod   PaymentMethod
	OrderItems      []OrderItem
	UserID          string
	User            User
	CreatedAt       *time.Time
}
