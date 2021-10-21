package product

import (
	"context"

	"github.com/vcnt72/try-golang-database-lib/entity"
)

type Repository interface {
	Store(ctx context.Context, product *entity.Product) (*entity.Product, error)
	Update(ctx context.Context, product *entity.Product) (*entity.Product, error)
	Delete(ctx context.Context, product *entity.Product) error
	FindByID(ctx context.Context, id string) (*entity.Product, error)
	Search(ctx context.Context, filterDTO FilterDTO) ([]entity.Product, error)
	Paginate(ctx context.Context, filterDTO FilterDTO, paginationOption *entity.Paginate) ([]entity.Product, error)
}

type Usecase interface {
	Store(ctx context.Context, createDTO CreateProductDTO) (*entity.Product, error)
	Update(ctx context.Context, updateDTO UpdateProductDTO) (*entity.Product, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*entity.Product, error)
	Paginate(ctx context.Context, filterDTO FilterDTO, paginationOption *entity.Paginate) ([]entity.Product, error)
}

type FilterDTO struct {
	ProductName string
	UserID      string
	ProductIDs  []string
}
