package postgres

import (
	"context"
	"inventory-service/domain"

	"inventory-service/model"

	"inventory-service/module/base/repository"

	"gorm.io/gorm"
)

type ProductsRepository interface {
	// Example operations
	repository.BaseRepository[domain.Products]
	ProductsList(
		ctx context.Context, tx *gorm.DB, page model.PaginationParam, order model.OrderParam, filter model.FilterParams,
		keyword model.KeywordParam,
	) (*model.PaginationData[domain.Products], error)
	FindProducts(
		ctx context.Context, tx *gorm.DB, id int64,
	) (*domain.Products, error)
}
