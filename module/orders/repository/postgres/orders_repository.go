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

type OrdersRepo struct {
	repository.BaseRepository[domain.Orders]
}

func NewOrdersRepository() OrdersRepository {
	keywordField := []string{
		"title",
	}

	repo := repository.NewBaseRepositoryImpl[domain.Orders](keywordField)
	return &OrdersRepo{
		BaseRepository: repo,
	}
}

func (r *OrdersRepo) OrdersList(
	ctx context.Context, tx *gorm.DB, page model.PaginationParam, order model.OrderParam, filter model.FilterParams,
	keyword model.KeywordParam,
) (*model.PaginationData[domain.Orders], error) {
	return r.Find(ctx, tx.Preload(clause.Associations), page, order, filter, keyword)
}

func (r *OrdersRepo) FindOrders(
	ctx context.Context, tx *gorm.DB, id int64,
) (*domain.Orders, error) {
	query := tx.WithContext(ctx).Preload(clause.Associations)

	result, err := r.FindByID(ctx, query, strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}

	return result, nil
}
