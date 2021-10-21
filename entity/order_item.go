package entity

type OrderItem struct {
	ID             string
	ProductID      string
	Product        Product
	Quantity       int
	OrderSummaryID string
	OrderSummary   OrderSummary
	ProductPrice   float64
}
