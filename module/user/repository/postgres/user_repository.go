package postgres

import (
	"inventory-service/domain"
	"inventory-service/module/base/repository"
)

type UserRepo struct {
	repository.BaseRepository[domain.User]
}

func NewUserRepository() UserRepository {
	keywordField := []string{
		"name",
	}

	repo := repository.NewBaseRepositoryImpl[domain.User](keywordField)
	return &UserRepo{
		BaseRepository: repo,
	}
}
