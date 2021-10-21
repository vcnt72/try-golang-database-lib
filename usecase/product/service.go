package product

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vcnt72/try-golang-database-lib/entity"
	"github.com/vcnt72/try-golang-database-lib/usecase/user"
	"go.uber.org/zap"
)

type service struct {
	repo        Repository
	userService user.Usecase
	logger      zap.Logger
}

func NewService(productRepository Repository, userService user.Usecase) Usecase {
	return &service{
		repo:        productRepository,
		userService: userService,
	}
}
func (s *service) Store(ctx context.Context, createDTO CreateProductDTO) (*entity.Product, error) {

	product := &entity.Product{
		Name:     createDTO.Name,
		Quantity: createDTO.Quantity,
		Price:    createDTO.Price,
	}

	product, err := s.repo.Store(ctx, product)
	if err != nil {
		s.logger.Sugar().Error(err)
		return nil, err
	}

	return product, nil

}

func (s *service) Update(ctx context.Context, updateDTO UpdateProductDTO) (*entity.Product, error) {
	product, err := s.GetByID(ctx, updateDTO.ID)

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

func (s *service) Delete(ctx context.Context, id string) error {

	product, err := s.GetByID(ctx, id)

	if err != nil {
		err = errors.Wrap(err, "error on get product by id")
		s.logger.Sugar().Error(err)
		return err
	}
	if err := s.repo.Delete(ctx, product); err != nil {
		s.logger.Sugar().Error(err)
		return err
	}

	return nil
}

func (s *service) GetByID(ctx context.Context, id string) (*entity.Product, error) {
	product, err := s.repo.FindByID(ctx, id)

	if err != nil {
		s.logger.Sugar().Error(err)
		return nil, err
	}

	s.logger.Sugar().Info("Success productService.GetByID")

	return product, nil
}

func (s *service) Paginate(ctx context.Context, filterDTO FilterDTO, paginationOption *entity.Paginate) ([]entity.Product, error) {
	products, err := s.repo.Paginate(ctx, filterDTO, paginationOption)

	if err != nil {
		return nil, err
	}

	return products, nil
}
