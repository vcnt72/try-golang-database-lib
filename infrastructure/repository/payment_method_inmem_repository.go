package repository

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/vcnt72/try-golang-database-lib/entity"
	"github.com/vcnt72/try-golang-database-lib/usecase/payment_method"
)

type paymentMethodInMemRepository struct {
	paymentMethods []*entity.PaymentMethod
}

func (r *paymentMethodInMemRepository) FindAll(ctx context.Context) ([]entity.PaymentMethod, error) {
	var paymentMethods []entity.PaymentMethod

	for _, paymentMethod := range r.paymentMethods {
		paymentMethods = append(paymentMethods, *paymentMethod)
	}

	return paymentMethods, nil
}

func (r *paymentMethodInMemRepository) FindByID(ctx context.Context, id string) (*entity.PaymentMethod, error) {
	for _, paymentMethod := range r.paymentMethods {
		if paymentMethod.ID == id {
			return paymentMethod, nil
		}
	}

	return nil, entity.ErrNotFound
}

func NewPaymentMethodInMemRepository() payment_method.Repository {
	return &paymentMethodInMemRepository{
		paymentMethods: []*entity.PaymentMethod{
			{
				ID:   uuid.NewV4().String(),
				Name: "Cash",
			},
			{
				ID:   uuid.NewV4().String(),
				Name: "E-Money",
			},
		},
	}
}
