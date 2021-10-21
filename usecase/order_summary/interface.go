package order_summary

import (
	"context"
	"time"

	"github.com/vcnt72/try-golang-database-lib/entity"
)

type Repository interface {
	Store(ctx context.Context, orderSummary *entity.OrderSummary) (*entity.OrderSummary, error)
	Search(ctx context.Context, filterDTO FilterDTO) ([]entity.OrderSummary, error)
}

type Usecase interface {
	Store(ctx context.Context, createDTO CreateOrderSummaryDTO) (*entity.OrderSummary, error)
	Search(ctx context.Context, filterDTO FilterDTO) ([]entity.OrderSummary, error)
}

type FilterDTO struct {
	StartDate time.Time
	EndDate   time.Time
}
