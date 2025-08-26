package postgres

// import (
// 	"context"

// 	"gitlab.com/integra_sm/cherry-v2-core-service/domain"
// 	"gitlab.com/integra_sm/cherry-v2-core-service/model"
// 	"gitlab.com/integra_sm/cherry-v2-core-service/module/base/repository"
// 	"gorm.io/gorm"
// )

// type AttributeValueRepository interface {
// 	// Example operations
// 	repository.BaseRepository[domain.AttributeValue]
// 	AttributeValueList(
// 		ctx context.Context, tx *gorm.DB, page model.PaginationParam, order model.OrderParam, filter model.FilterParams,
// 		keyword model.KeywordParam,
// 	) (*model.PaginationData[domain.AttributeValue], error)
// 	FindAttributeValue(
// 		ctx context.Context, tx *gorm.DB, refId string,
// 	) (*domain.AttributeValue, error)
// }

// type AttributeRepository interface {
// 	// Example operations
// 	repository.BaseRepository[domain.Attribute]
// 	AttributeList(
// 		ctx context.Context, tx *gorm.DB, page model.PaginationParam, order model.OrderParam, filter model.FilterParams,
// 		keyword model.KeywordParam) (*model.PaginationData[domain.Attribute], error)
// }
