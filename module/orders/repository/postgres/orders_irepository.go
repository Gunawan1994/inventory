package postgres

import (
	"context"
	"inventory-service/domain"

	"inventory-service/model"

	"inventory-service/module/base/repository"

	"gorm.io/gorm"
)

type OrdersRepository interface {
	// Example operations
	repository.BaseRepository[domain.Orders]
	OrdersList(
		ctx context.Context, tx *gorm.DB, page model.PaginationParam, order model.OrderParam, filter model.FilterParams,
		keyword model.KeywordParam,
	) (*model.PaginationData[domain.Orders], error)
	FindOrders(
		ctx context.Context, tx *gorm.DB, id int64,
	) (*domain.Orders, error)
}
