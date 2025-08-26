package postgres

import (
	"inventory-service/domain"
	"inventory-service/module/base/repository"
)

type UserRepository interface {
	repository.BaseRepository[domain.User]
}
