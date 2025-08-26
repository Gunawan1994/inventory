package usecase

import (
	"context"

	"inventory-service/model"
)

type OrdersUseCase interface {
	Create(ctx context.Context, req *model.CreateOrdersReq) (*model.CreateOrdersRes, error)
	GetById(ctx context.Context, req *model.GetIdOrdersReq) (*model.GetIdOrdersRes, error)
	GetList(ctx context.Context, req *model.GetListOrdersReq) (
		*model.GetListOrdersRes, error,
	)
	Update(ctx context.Context, req *model.UpdateOrdersReq) (*model.UpdateOrdersRes, error)
	Delete(ctx context.Context, req *model.DeleteOrdersReq) (*model.DeleteOrdersRes, error)
}
