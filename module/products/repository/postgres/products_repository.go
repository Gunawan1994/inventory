package postgres

import (
	"context"
	"strconv"

	"inventory-service/domain"
	"inventory-service/model"
	"inventory-service/module/base/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductsRepo struct {
	repository.BaseRepository[domain.Products]
}

func NewProductsRepository() ProductsRepository {
	keywordField := []string{
		"title",
	}

	repo := repository.NewBaseRepositoryImpl[domain.Products](keywordField)
	return &ProductsRepo{
		BaseRepository: repo,
	}
}

func (r *ProductsRepo) ProductsList(
	ctx context.Context, tx *gorm.DB, page model.PaginationParam, order model.OrderParam, filter model.FilterParams,
	keyword model.KeywordParam,
) (*model.PaginationData[domain.Products], error) {
	return r.Find(ctx, tx.Preload(clause.Associations), page, order, filter, keyword)
}

func (r *ProductsRepo) FindProducts(
	ctx context.Context, tx *gorm.DB, id int64,
) (*domain.Products, error) {
	query := tx.WithContext(ctx).Preload(clause.Associations)

	result, err := r.FindByID(ctx, query, strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}

	return result, nil
}
