package usecase

import (
	"context"

	"inventory-service/domain"
	"inventory-service/model"
)

type AuthUsecase interface {
	VerifyCredential(ctx context.Context, req model.VerifyCredential) (*domain.User, error)
	RegisterUser(ctx context.Context, req *model.CreateUserReq) (*domain.User, error)
}
