package usecase

import (
	"context"

	"inventory-service/model"
	"inventory-service/module/user/repository/postgres"

	"gorm.io/gorm"
)

type UserUseCaseImpl struct {
	db   *gorm.DB
	repo postgres.UserRepository
}

func NewUserUseCase(
	db *gorm.DB, repo postgres.UserRepository,
) UserUsecase {
	return &UserUseCaseImpl{
		db:   db,
		repo: repo,
	}
}

func (s *UserUseCaseImpl) Create(
	ctx context.Context, req *model.CreateUserReq,
) (*model.CreateUserRes, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	body := req.ToDomain()

	if err := s.repo.CreateTx(ctx, tx, body); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &model.CreateUserRes{User: body}, nil
}

func (s *UserUseCaseImpl) GetUserByEmail(
	ctx context.Context, email string,
) (*model.User, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	user, err := s.repo.FindUserByEmail(ctx, tx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &model.User{
		Id:        user.Id,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
