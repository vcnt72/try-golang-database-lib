package order_summary

import (
	"context"
	"errors"

	"github.com/vcnt72/try-golang-database-lib/entity"
	"github.com/vcnt72/try-golang-database-lib/usecase/payment_method"
	"github.com/vcnt72/try-golang-database-lib/usecase/product"
	"github.com/vcnt72/try-golang-database-lib/usecase/user"
	"go.uber.org/zap"
)

type service struct {
	orderSummaryRepository  Repository
	userRepository          user.Repository
	productRepository       product.Repository
	paymentMethodRepository payment_method.Repository
	logger                  *zap.Logger
}

func NewService(orderSummaryRepository Repository, userRepository user.Repository, productRepository product.Repository, paymentMethodRepository payment_method.Repository, logger *zap.Logger) Usecase {
	return &service{
		orderSummaryRepository:  orderSummaryRepository,
		userRepository:          userRepository,
		productRepository:       productRepository,
		paymentMethodRepository: paymentMethodRepository,
		logger:                  logger,
	}
}
func (s *service) Store(ctx context.Context, createDTO CreateOrderSummaryDTO) (*entity.OrderSummary, error) {
	paymentMethod, err := s.paymentMethodRepository.FindByID(ctx, createDTO.PaymentMethodID)

	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindByID(ctx, createDTO.UserID)

	if err != nil {
		return nil, err
	}

	orderItems, err := s.genOrderItems(ctx, createDTO.OrderItems)

	if err != nil {
		return nil, err
	}

	orderSummary := &entity.OrderSummary{
		PaymentMethod: *paymentMethod,
		User:          *user,
		OrderItems:    orderItems,
	}

	orderSummary, err = s.orderSummaryRepository.Store(ctx, orderSummary)

	if err != nil {
		return nil, err
	}

	return orderSummary, nil
}

func (s *service) genOrderItems(ctx context.Context, createDTO []CreateOrderItemDTO) ([]entity.OrderItem, error) {

	var orderItems []entity.OrderItem

	for _, v := range createDTO {

		product, err := s.productRepository.FindByID(ctx, v.ProductID)

		if errors.Is(err, entity.ErrNotFound) {
			return nil, ErrProductMissing
		}

		if err != nil {
			return nil, err
		}

		orderItem := &entity.OrderItem{
			Product:      *product,
			Quantity:     v.Quantity,
			ProductPrice: product.Price,
		}

		orderItems = append(orderItems, *orderItem)
	}

	return orderItems, nil
}

func (s *service) Search(ctx context.Context, filterDTO FilterDTO) ([]entity.OrderSummary, error) {
	orderSummaries, err := s.orderSummaryRepository.Search(ctx, filterDTO)

	if err != nil {
		return nil, err
	}

	return orderSummaries, nil
}
