package repository

import (
	"context"

	"github.com/vcnt72/try-golang-database-lib/entity"
	"github.com/vcnt72/try-golang-database-lib/usecase/product"
)

type productInMemRepository struct {
	products []*entity.Product
}

func (r *productInMemRepository) Store(ctx context.Context, product *entity.Product) (*entity.Product, error) {

	r.products = append(r.products, product)
	return product, nil
}

func (r *productInMemRepository) Update(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	for i, val := range r.products {
		if val.ID == product.ID {
			r.products[i] = product
		}
	}

	return product, nil
}

func (r *productInMemRepository) Delete(ctx context.Context, product *entity.Product) error {

	targetIndex := -1

	for i, val := range r.products {
		if val.ID == product.ID {
			targetIndex = i
			break
		}
	}

	if targetIndex == -1 {
		return entity.ErrNotFound
	}

	r.products[targetIndex] = r.products[len(r.products)-1]
	r.products[len(r.products)-1] = nil
	r.products = r.products[:len(r.products)-1]

	return nil
}

func (r *productInMemRepository) FindByID(ctx context.Context, id string) (*entity.Product, error) {
	for _, val := range r.products {
		if val.ID == id {
			return val, nil
			break
		}
	}
	return nil, entity.ErrNotFound
}

func (r *productInMemRepository) Search(ctx context.Context, filterDTO product.FilterDTO) ([]entity.Product, error) {
	// I'am lazy to create the logic. Maybe later
	var products []entity.Product

	for _, v := range r.products {
		products = append(products, *v)
	}

	return products, nil
}

func (r *productInMemRepository) Paginate(ctx context.Context, filterDTO product.FilterDTO, paginationOption *entity.Paginate) ([]entity.Product, error) {
	// I'am lazy to create the logic. Maybe later
	var products []entity.Product

	for _, v := range r.products {
		products = append(products, *v)
	}

	return products, nil
}

func NewProductInMemRepository() product.Repository {
	return &productInMemRepository{}
}
