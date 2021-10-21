package payment_method

import (
	"context"

	"github.com/vcnt72/try-golang-database-lib/entity"
	"go.uber.org/zap"
)

type service struct {
	repo   Repository
	logger zap.Logger
}

func (s *service) GetAll(ctx context.Context) ([]entity.PaymentMethod, error) {
	paymentMethods, err := s.repo.FindAll(ctx)

	if err != nil {
		s.logger.Sugar().Error(err.Error())
		return nil, err
	}
	s.logger.Sugar().Info("Success running paymentMethodService.GetAll")
	return paymentMethods, nil
}

func (s *service) GetByID(ctx context.Context, id string) (*entity.PaymentMethod, error) {
	paymentMethod, err := s.repo.FindByID(ctx, id)

	if err != nil {
		s.logger.Sugar().Error(err.Error())
		return nil, err
	}
	s.logger.Sugar().Info("Success running paymentMethodService.GetByID")

	return paymentMethod, nil
}

func NewService(paymentMethodRepository Repository) Usecase {
	return &service{
		repo: paymentMethodRepository,
	}
}
