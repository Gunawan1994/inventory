package usecase

import (
	"context"

	"inventory-service/model"
)

type ProductsUseCase interface {
	Create(ctx context.Context, req *model.CreateProductsReq) (*model.CreateProductsRes, error)
	GetById(ctx context.Context, req *model.GetIdProductsReq) (*model.GetIdProductsRes, error)
	GetList(ctx context.Context, req *model.GetListProductsReq) (
		*model.GetListProductsRes, error,
	)
	Update(ctx context.Context, req *model.UpdateProductsReq) (*model.UpdateProductsRes, error)
	Delete(ctx context.Context, req *model.DeleteProductsReq) (*model.DeleteProductsRes, error)
}
