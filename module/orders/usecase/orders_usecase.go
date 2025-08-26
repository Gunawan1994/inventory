package usecase

import (
	"context"
	"errors"
	"strconv"

	"inventory-service/model"
	"inventory-service/module/orders/repository/postgres"

	"gorm.io/gorm"
)

type OrdersUseCaseImpl struct {
	db   *gorm.DB
	repo postgres.OrdersRepository
}

func NewOrdersUseCase(
	db *gorm.DB, repo postgres.OrdersRepository,
) OrdersUseCase {
	return &OrdersUseCaseImpl{
		db:   db,
		repo: repo,
	}
}

func (s *OrdersUseCaseImpl) Create(
	ctx context.Context, req *model.CreateOrdersReq,
) (*model.CreateOrdersRes, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	body := req.ToDomain()

	if err := s.repo.CreateTx(ctx, tx, body); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &model.CreateOrdersRes{Orders: body}, nil
}

func (s *OrdersUseCaseImpl) GetById(
	ctx context.Context, req *model.GetIdOrdersReq,
) (*model.GetIdOrdersRes, error) {
	result, err := s.repo.FindByID(ctx, s.db, strconv.FormatInt(req.Id, 10))
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("Orders not found")
	}

	return &model.GetIdOrdersRes{Orders: result}, nil
}

func (s *OrdersUseCaseImpl) Update(
	ctx context.Context, req *model.UpdateOrdersReq,
) (*model.UpdateOrdersRes, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	body := req.ToDomain()

	data, err := s.repo.FindByID(ctx, s.db, strconv.FormatInt(req.Id, 10))
	if err != nil {
		return nil, err
	}

	body.Id = data.Id

	if err := s.repo.UpdateTx(ctx, tx, body); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &model.UpdateOrdersRes{Orders: body}, nil

}

func (s *OrdersUseCaseImpl) GetList(
	ctx context.Context, req *model.GetListOrdersReq,
) (*model.GetListOrdersRes, error) {
	result, err := s.repo.OrdersList(ctx, s.db, req.Page, req.Order, req.Filter, req.Keyword)
	if err != nil {
		return nil, err
	}

	return &model.GetListOrdersRes{
		Data:       result.Data,
		Pagination: result.ToPagination(),
	}, nil
}

func (s *OrdersUseCaseImpl) Delete(
	ctx context.Context, req *model.DeleteOrdersReq,
) (*model.DeleteOrdersRes, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	data, err := s.repo.FindByID(ctx, s.db, strconv.FormatInt(req.Id, 10))
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("Orders not found")
	}

	if err := s.repo.DeleteByIDTx(ctx, tx, strconv.FormatInt(req.Id, 10)); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &model.DeleteOrdersRes{Orders: data}, nil
}
