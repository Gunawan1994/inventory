package postgres

// import (
// 	"context"

// 	"gitlab.com/integra_sm/cherry-v2-core-service/domain"
// 	"gitlab.com/integra_sm/cherry-v2-core-service/model"
// 	"gitlab.com/integra_sm/cherry-v2-core-service/module/base/repository"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/clause"
// )

// type AttributeValueRepo struct {
// 	repository.BaseRepository[domain.AttributeValue]
// }

// func NewAttributeValueRepository() AttributeValueRepository {
// 	keywordField := []string{
// 		"name",
// 	}

// 	repo := repository.NewBaseRepositoryImpl[domain.AttributeValue](keywordField, nil)
// 	return &AttributeValueRepo{
// 		BaseRepository: repo,
// 	}
// }

// func (r *AttributeValueRepo) AttributeValueList(
// 	ctx context.Context, tx *gorm.DB, page model.PaginationParam, order model.OrderParam, filter model.FilterParams,
// 	keyword model.KeywordParam,
// ) (*model.PaginationData[domain.AttributeValue], error) {
// 	return r.Find(ctx, tx.Preload(clause.Associations), page, order, filter, keyword)
// }

// func (r *AttributeValueRepo) FindAttributeValue(
// 	ctx context.Context, tx *gorm.DB, refId string,
// ) (*domain.AttributeValue, error) {
// 	query := tx.WithContext(ctx).Preload(clause.Associations)

// 	var attribute domain.AttributeValue
// 	err := query.Model(&domain.AttributeValue{}).Where("references_id = ?", refId).First(&attribute).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &attribute, nil
// }
