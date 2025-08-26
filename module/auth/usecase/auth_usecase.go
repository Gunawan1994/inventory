package usecase

import (
	"context"
	"inventory-service/helpers/utils"
	"inventory-service/model"
	"inventory-service/module/auth/repository/postgres"
	userRepo "inventory-service/module/user/repository/postgres"

	"inventory-service/domain"

	"gorm.io/gorm"
)

func NewAuthUseCase(
	db *gorm.DB, repo postgres.AuthIRepository, userRepo userRepo.UserRepository,
) AuthUsecase {
	return &AuthUsecaseImpl{
		db:       db,
		repo:     repo,
		repoUser: userRepo,
	}
}

type AuthUsecaseImpl struct {
	db       *gorm.DB
	repo     postgres.AuthIRepository
	repoUser userRepo.UserRepository
}

func (a *AuthUsecaseImpl) VerifyCredential(ctx context.Context, req model.VerifyCredential) (*domain.User, error) {
	tx := a.db.Begin()
	defer tx.Rollback()

	result, err := a.repo.VerifyCredential(ctx, tx, req)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s *AuthUsecaseImpl) RegisterUser(
	ctx context.Context, req *model.CreateUserReq,
) (*domain.User, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	body := req.ToDomain()

	if err := s.repoUser.CreateTx(ctx, tx, &domain.User{
		Id:       body.Id,
		Email:    body.Email,
		Password: utils.HashValue(body.Password),
		Username: body.Username,
	}); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &domain.User{
		Id:       body.Id,
		Email:    body.Email,
		Username: body.Username,
	}, nil
}
