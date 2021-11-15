package payment_method

import (
	"context"
	"errors"

	"github.com/vcnt72/try-golang-database-lib/entity"
	"go.uber.org/zap"
)

type service struct {
	paymentMethodRepository Repository
	logger                  *zap.Logger
}

func (s *service) GetAll(ctx context.Context) ([]entity.PaymentMethod, error) {
	paymentMethods, err := s.paymentMethodRepository.FindAll(ctx)

	if err != nil {
		s.logger.Sugar().Error(err.Error())
		return nil, err
	}
	s.logger.Sugar().Info("Success running paymentMethodService.GetAll")
	return paymentMethods, nil
}

func (s *service) GetByID(ctx context.Context, id string) (*entity.PaymentMethod, error) {
	paymentMethod, err := s.paymentMethodRepository.FindByID(ctx, id)

	if errors.Is(err, entity.ErrNotFound) {
		return nil, err
	}

	if err != nil {
		s.logger.Sugar().Error(err.Error())
		return nil, err
	}

	s.logger.Sugar().Info("Success running paymentMethodService.GetByID")

	return paymentMethod, nil
}

func NewService(paymentMethodRepository Repository, logger *zap.Logger) Usecase {
	return &service{
		paymentMethodRepository: paymentMethodRepository,
		logger:                  logger,
	}
}
