package payment_method

import (
	"context"

	"github.com/vcnt72/try-golang-database-lib/entity"
)

type Repository interface {
	FindAll(ctx context.Context) ([]entity.PaymentMethod, error)
	FindByID(ctx context.Context, id string) (*entity.PaymentMethod, error)
}

type Usecase interface {
	GetAll(ctx context.Context) ([]entity.PaymentMethod, error)
	GetByID(ctx context.Context, id string) (*entity.PaymentMethod, error)
}
