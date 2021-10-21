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
	repo                 Repository
	userService          user.Usecase
	productService       product.Usecase
	paymentMethodService payment_method.Usecase
	logger               zap.Logger
}

func NewService(repo Repository, userService user.Usecase, productService product.Usecase, paymentMethodService payment_method.Usecase, logger zap.Logger) Usecase {
	return &service{
		repo:                 repo,
		userService:          userService,
		productService:       productService,
		paymentMethodService: paymentMethodService,
		logger:               logger,
	}
}
func (s *service) Store(ctx context.Context, createDTO CreateOrderSummaryDTO) (*entity.OrderSummary, error) {
	paymentMethod, err := s.paymentMethodService.GetByID(ctx, createDTO.PaymentMethodID)

	if err != nil {
		return nil, err
	}

	user, err := s.userService.GetByID(ctx, createDTO.UserID)

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

	orderSummary, err = s.repo.Store(ctx, orderSummary)

	if err != nil {
		return nil, err
	}

	return orderSummary, nil
}

func (s *service) genOrderItems(ctx context.Context, createDTO []CreateOrderItemDTO) ([]entity.OrderItem, error) {

	var orderItems []entity.OrderItem

	for _, v := range createDTO {

		product, err := s.productService.GetByID(ctx, v.ProductID)

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
	orderSummaries, err := s.repo.Search(ctx, filterDTO)

	if err != nil {
		return nil, err
	}

	return orderSummaries, nil
}
