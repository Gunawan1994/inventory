package postgres

// import (
// 	"context"

// 	"gitlab.com/integra_sm/cherry-v2-core-service/domain"
// 	"gitlab.com/integra_sm/cherry-v2-core-service/model"
// 	"gitlab.com/integra_sm/cherry-v2-core-service/module/base/repository"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/clause"
// )

// type AttributeRepo struct {
// 	repository.BaseRepository[domain.Attribute]
// }

// func NewAttributeRepository() AttributeRepository {
// 	keywordField := []string{
// 		"name",
// 	}

// 	repo := repository.NewBaseRepositoryImpl[domain.Attribute](keywordField, nil)
// 	return &AttributeRepo{
// 		BaseRepository: repo,
// 	}
// }

// func (r *AttributeRepo) AttributeList(
// 	ctx context.Context, tx *gorm.DB, page model.PaginationParam, order model.OrderParam, filter model.FilterParams,
// 	keyword model.KeywordParam,
// ) (*model.PaginationData[domain.Attribute], error) {
// 	return r.Find(ctx, tx.Preload(clause.Associations), page, order, filter, keyword)
// }
