package repository

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/vcnt72/try-golang-database-lib/entity"
	"github.com/vcnt72/try-golang-database-lib/usecase/order_summary"
)

type orderSummaryInMemRepository struct {
	orderSummaries []*entity.OrderSummary
}

func (r *orderSummaryInMemRepository) Store(ctx context.Context, orderSummary *entity.OrderSummary) (*entity.OrderSummary, error) {

	now := time.Now().UTC()
	orderSummary.ID = uuid.NewV4().String()
	orderSummary.CreatedAt = &now

	r.orderSummaries = append(r.orderSummaries, orderSummary)
	return orderSummary, nil
}

func (r *orderSummaryInMemRepository) Search(ctx context.Context, filterDTO order_summary.FilterDTO) ([]entity.OrderSummary, error) {

	var orderSummaries []entity.OrderSummary
	for _, orderSummary := range r.orderSummaries {
		if orderSummary.CreatedAt.After(filterDTO.StartDate) && orderSummary.CreatedAt.Before(filterDTO.EndDate) {
			orderSummaries = append(orderSummaries, *orderSummary)
		}
	}

	return orderSummaries, nil
}

func NewOrderSummaryInMemRepository() order_summary.Repository {
	return &orderSummaryInMemRepository{}
}
