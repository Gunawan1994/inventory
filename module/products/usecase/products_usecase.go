package usecase

import (
	"context"
	"errors"
	"strconv"

	"inventory-service/model"
	"inventory-service/module/products/repository/postgres"

	"gorm.io/gorm"
)

type ProductsUseCaseImpl struct {
	db   *gorm.DB
	repo postgres.ProductsRepository
}

func NewProductsUseCase(
	db *gorm.DB, repo postgres.ProductsRepository,
) ProductsUseCase {
	return &ProductsUseCaseImpl{
		db:   db,
		repo: repo,
	}
}

func (s *ProductsUseCaseImpl) Create(
	ctx context.Context, req *model.CreateProductsReq,
) (*model.CreateProductsRes, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	body := req.ToDomain()

	if err := s.repo.CreateTx(ctx, tx, body); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &model.CreateProductsRes{Products: body}, nil
}

func (s *ProductsUseCaseImpl) GetById(
	ctx context.Context, req *model.GetIdProductsReq,
) (*model.GetIdProductsRes, error) {
	result, err := s.repo.FindByID(ctx, s.db, strconv.FormatInt(req.Id, 10))
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("Products not found")
	}

	return &model.GetIdProductsRes{Products: result}, nil
}

func (s *ProductsUseCaseImpl) Update(
	ctx context.Context, req *model.UpdateProductsReq,
) (*model.UpdateProductsRes, error) {
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
	return &model.UpdateProductsRes{Products: body}, nil

}

func (s *ProductsUseCaseImpl) GetList(
	ctx context.Context, req *model.GetListProductsReq,
) (*model.GetListProductsRes, error) {
	result, err := s.repo.ProductsList(ctx, s.db, req.Page, req.Order, req.Filter, req.Keyword)
	if err != nil {
		return nil, err
	}

	return &model.GetListProductsRes{
		Data:       result.Data,
		Pagination: result.ToPagination(),
	}, nil
}

func (s *ProductsUseCaseImpl) Delete(
	ctx context.Context, req *model.DeleteProductsReq,
) (*model.DeleteProductsRes, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	data, err := s.repo.FindByID(ctx, s.db, strconv.FormatInt(req.Id, 10))
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("Products not found")
	}

	if err := s.repo.DeleteByIDTx(ctx, tx, strconv.FormatInt(req.Id, 10)); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &model.DeleteProductsRes{Products: data}, nil
}
