package product

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vcnt72/try-golang-database-lib/entity"
	"github.com/vcnt72/try-golang-database-lib/usecase/user"
	"go.uber.org/zap"
)

type Service struct {
	productRepository Repository
	userService       user.Usecase
	logger            *zap.Logger
}

func NewService(productRepository Repository, userService user.Usecase, logger *zap.Logger) Usecase {
	return &Service{
		productRepository: productRepository,
		userService:       userService,
		logger:            logger,
	}
}
func (s *Service) Store(ctx context.Context, createDTO CreateProductDTO) (*entity.Product, error) {

	product := &entity.Product{
		Name:     createDTO.Name,
		Quantity: createDTO.Quantity,
		Price:    createDTO.Price,
	}

	product, err := s.productRepository.Store(ctx, product)

	if err != nil {
		s.logger.Sugar().Error(err)
		return nil, err
	}

	return product, nil

}

func (s *Service) Update(ctx context.Context, updateDTO UpdateProductDTO) (*entity.Product, error) {
	product, err := s.GetByID(ctx, updateDTO.ID)

	if errors.Is(err, entity.ErrNotFound) {
		return nil, err
	}

	if err != nil {
		err = errors.Wrap(err, "error on get product by id")
		s.logger.Sugar().Error(err)
		return nil, err
	}

	product.Name = updateDTO.Name
	product.Price = updateDTO.Price
	product.Quantity = updateDTO.Quantity

	return product, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {

	product, err := s.productRepository.FindByID(ctx, id)

	if errors.Is(err, entity.ErrNotFound) {
		return err
	}

	if err != nil {
		err = errors.Wrap(err, "error on get product by id")
		s.logger.Sugar().Error(err)
		return err
	}
	if err := s.productRepository.Delete(ctx, product); err != nil {
		s.logger.Sugar().Error(err)
		return err
	}

	return nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*entity.Product, error) {
	product, err := s.productRepository.FindByID(ctx, id)

	if err != nil {
		s.logger.Sugar().Error(err)
		return nil, err
	}

	s.logger.Sugar().Info("Success productService.GetByID")

	return product, nil
}

func (s *Service) Paginate(ctx context.Context, filterDTO FilterDTO, paginationOption *entity.Paginate) ([]entity.Product, error) {
	products, err := s.productRepository.Paginate(ctx, filterDTO, paginationOption)

	if err != nil {
		return nil, err
	}

	return products, nil
}
