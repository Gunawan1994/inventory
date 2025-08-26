package postgres

import (
	"context"
	"inventory-service/model"

	"inventory-service/domain"

	"gorm.io/gorm"
)

type AuthIRepository interface {
	// repository.BaseRepository[domain.User]
	VerifyCredential(ctx context.Context, tx *gorm.DB, req model.VerifyCredential) (*domain.User, error)
}
