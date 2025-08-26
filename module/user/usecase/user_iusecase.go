package usecase

import (
	"context"
	"inventory-service/model"
)

type UserUsecase interface {
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, req *model.CreateUserReq) (*model.CreateUserRes, error)
}
