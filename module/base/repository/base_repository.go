package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"reflect"
	"strings"

	"inventory-service/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepository[T any] interface {
	CreateTx(ctx context.Context, tx *gorm.DB, data *T) error
	CreateUnscoped(ctx context.Context, tx *gorm.DB, data *T) error
	UpdateAssociationMany2ManyTx(tx *gorm.DB, data *T) error
	UpdateTx(ctx context.Context, tx *gorm.DB, data *T) error
	UpdateTxWithAssociations(ctx context.Context, tx *gorm.DB, data *T) error
	DeleteByIDTx(ctx context.Context, tx *gorm.DB, id string) error
	Delete(ctx context.Context, tx *gorm.DB, column, id string) error
	Find(
		ctx context.Context, tx *gorm.DB, page model.PaginationParam, order model.OrderParam, filter model.FilterParams,
		keyword model.KeywordParam,
	) (*model.PaginationData[T], error)
	FindByPagination(
		ctx context.Context, tx *gorm.DB, page model.PaginationParam,
	) (*model.PaginationData[T], error)
	FindWithQuery(
		ctx context.Context, tx *gorm.DB, order model.OrderParam, filter model.FilterParams,
	) (*[]T, error)
	FindByID(ctx context.Context, tx *gorm.DB, id string) (*T, error)
	PaginationQuery(page, pageSize int, query *gorm.DB) (PaginationResult[T], error)
	OrderQuery(param model.OrderParam, query *gorm.DB) *gorm.DB
	FilterQuery(filter model.FilterParams, query *gorm.DB) *gorm.DB
	FindUserByEmail(
		ctx context.Context, tx *gorm.DB, email string,
	) (*T, error)
}

type BaseRepositoryImpl[T any] struct {
	keywordFields []string
}

func NewBaseRepositoryImpl[T any](
	keywordFields []string,
) BaseRepository[T] {
	return &BaseRepositoryImpl[T]{
		keywordFields: keywordFields,
	}
}

type PaginationResult[T any] struct {
	Page             int   // The current page
	PageSize         int   // The size of the page
	TotalPage        int64 // The total number of pages
	TotalDataPerPage int64 // The total number of data per page
	TotalData        int64 // The total number of data
	Data             []*T  // The actual data
}

func (r *BaseRepositoryImpl[T]) CreateTx(ctx context.Context, tx *gorm.DB, data *T) error {
	if err := tx.WithContext(ctx).Omit(clause.Associations).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).
		Create(data).Error; err != nil {
		slog.Error("failed to create", slog.Any("error", err))
		return err
	}
	return nil
}

func (r *BaseRepositoryImpl[T]) CreateUnscoped(ctx context.Context, tx *gorm.DB, data *T) error {
	if err := tx.WithContext(ctx).Omit(clause.Associations).
		Create(data).Error; err != nil {
		slog.Error("failed to create", slog.Any("error", err))
		return err
	}
	return nil
}

func (r *BaseRepositoryImpl[T]) UpdateAssociationMany2ManyTx(tx *gorm.DB, data *T) error {
	val := reflect.ValueOf(data).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag.Get("gorm")

		if strings.Contains(tag, "many2many") {
			associationName := typeField.Name
			if err := tx.Model(data).Association(associationName).Replace(field.Interface()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *BaseRepositoryImpl[T]) UpdateTx(ctx context.Context, tx *gorm.DB, data *T) error {
	if err := tx.WithContext(ctx).Omit(clause.Associations).Model(data).Select("*").Updates(data).Error; err != nil {
		slog.Error("failed to update", slog.Any("error", err))
		return err
	}
	return nil
}

func (r *BaseRepositoryImpl[T]) UpdateTxWithAssociations(ctx context.Context, tx *gorm.DB, data *T) error {
	if err := tx.WithContext(ctx).Model(data).Select("*").Updates(data).Error; err != nil {
		slog.Error("failed to update", slog.Any("error", err))
		return err
	}
	return nil
}

func (r *BaseRepositoryImpl[T]) DeleteByIDTx(ctx context.Context, tx *gorm.DB, id string) error {
	if err := tx.WithContext(ctx).Unscoped().Where("id = ?", id).Delete(new(T)).Error; err != nil {
		slog.Error("failed to delete", slog.Any("error", err))
		return err
	}
	return nil
}

func (r *BaseRepositoryImpl[T]) Delete(ctx context.Context, tx *gorm.DB, column, id string) error {
	if err := tx.WithContext(ctx).Unscoped().Where(column+" = ?", id).Delete(new(T)).Error; err != nil {
		slog.Error("failed to delete", slog.Any("error", err))
		return err
	}
	return nil
}

func (r *BaseRepositoryImpl[T]) Find(
	ctx context.Context, tx *gorm.DB, page model.PaginationParam, order model.OrderParam, filter model.FilterParams,
	keyword model.KeywordParam,
) (*model.PaginationData[T], error) {
	query := tx.WithContext(ctx).Omit(clause.Associations)
	query = r.FilterQuery(filter, query)
	query = r.OrderQuery(order, query)
	query = r._keywordWhere(keyword, query)
	result, err := r.PaginationQuery(page.Offset, page.Limit, query)
	if err != nil {
		return nil, err
	}
	return &model.PaginationData[T]{
		Offset:          result.Page,
		Limit:           result.PageSize,
		TotalPages:      result.TotalPage,
		TotalRowPerPage: result.TotalDataPerPage,
		TotalRows:       result.TotalData,
		Data:            result.Data,
	}, nil
}

func (r *BaseRepositoryImpl[T]) FindByPagination(
	ctx context.Context, tx *gorm.DB, page model.PaginationParam,
) (*model.PaginationData[T], error) {
	query := tx.WithContext(ctx).Omit(clause.Associations)
	result, err := r.PaginationQuery(page.Offset, page.Limit, query)
	if err != nil {
		return nil, err
	}
	return &model.PaginationData[T]{
		Offset:          result.Page,
		Limit:           result.PageSize,
		TotalPages:      result.TotalPage,
		TotalRowPerPage: result.TotalDataPerPage,
		TotalRows:       result.TotalData,
		Data:            result.Data,
	}, nil
}
func (r *BaseRepositoryImpl[T]) FindWithQuery(
	ctx context.Context, tx *gorm.DB, order model.OrderParam, filter model.FilterParams,
) (*[]T, error) {
	var data *[]T
	query := tx.WithContext(ctx).Omit(clause.Associations)
	query = r.FilterQuery(filter, query)
	query = r.OrderQuery(order, query)
	if err := query.Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		slog.Error("failed to find all", err)
		return nil, err
	}
	return data, nil
}
func (r *BaseRepositoryImpl[T]) FindByID(ctx context.Context, tx *gorm.DB, id string) (*T, error) {
	var data T
	if err := tx.WithContext(ctx).Preload(clause.Associations).Where("id = ?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		slog.Error("failed to find by id", slog.Any("error", err))
		return nil, err
	}
	return &data, nil
}

func (r *BaseRepositoryImpl[T]) PaginationQuery(page, pageSize int, query *gorm.DB) (PaginationResult[T], error) {
	if pageSize < 0 {
		pageSize = -1
	}
	if page < 1 {
		page = -1
	}

	var total int64
	var data []*T
	if err := query.Model(new(T)).Count(&total).Error; err != nil {
		return PaginationResult[T]{}, err
	}
	if err := query.Offset(page).Limit(pageSize).Find(&data).Error; err != nil {
		return PaginationResult[T]{}, err
	}

	var totalPage int64
	if pageSize > 0 {
		totalPage = total / int64(pageSize)
		if total%int64(pageSize) > 0 {
			totalPage++
		}
	} else if pageSize == 0 {
		totalPage = 0
	} else {
		totalPage = 1
	}

	return PaginationResult[T]{
		Page:             page,
		PageSize:         pageSize,
		TotalPage:        totalPage,
		TotalDataPerPage: int64(len(data)),
		TotalData:        total,
		Data:             data,
	}, nil
}

func (r *BaseRepositoryImpl[T]) _keywordWhere(keyword model.KeywordParam, query *gorm.DB) *gorm.DB {
	if keyword.Value != "" {
		keyword.Value = "%" + strings.ToLower(keyword.Value) + "%"
		//for i, field := range r.keywordFields {
		//	field = "lower(" + field + ")"
		//	if i == 0 {
		//		query = query.Where(field+" like ?", keyword.Value)
		//	} else {
		//		query = query.Or(field+" like ?", keyword.Value)
		//	}
		//}
		for _, field := range r.keywordFields {
			field = "lower(" + field + ")"
			//if i == 0 {
			//	query = query.Where(field+" like ?", keyword.Value)
			//} else {
			query = query.Or(field+" like ?", keyword.Value)
			//}
		}
	}
	return query
}

func (r *BaseRepositoryImpl[T]) FilterQuery(filter model.FilterParams, query *gorm.DB) *gorm.DB {
	for _, f := range filter {
		keylist := r._generateWhere(*f)
		if strings.Contains(f.Field, "json") {
			switch query.Dialector.Name() {
			case "postgres":
				f.Field = strings.Replace(f.Field, "json_", "", -1) + "::text"
			}
		}
		switch f.Operator {
		case "like":
			f.Field = "lower(" + f.Field + ")"
			query = query.Where(fmt.Sprintf("%s %s ?", f.Field, f.Operator), keylist)
		case "in", "not in":
			query = query.Where(fmt.Sprintf("%s %s (?)", f.Field, f.Operator), keylist)
		default:
			query = query.Where(fmt.Sprintf("%s %s ?", f.Field, f.Operator), f.Value)
		}
	}
	return query
}

func (r *BaseRepositoryImpl[T]) OrderQuery(param model.OrderParam, query *gorm.DB) *gorm.DB {
	if param.Order != "" && param.OrderBy != "" {
		query = query.Order(fmt.Sprintf("%s %s", param.OrderBy, param.Order))
	}
	return query
}

func (r *BaseRepositoryImpl[T]) _generateWhere(filter model.FilterParam) []interface{} {
	keySearch := strings.ToLower(filter.Value)
	var keyList []interface{}

	if filter.Operator == "like" {
		keyList = make([]interface{}, 1)
		keyList[0] = "%" + keySearch + "%"
	} else if filter.Operator == "in" || filter.Operator == "not in" {
		keys := strings.Split(keySearch, ",")
		keyList = make([]interface{}, len(keys))
		for i, key := range keys {
			keyList[i] = key
		}
	} else {
		keyList = make([]interface{}, 1)
		keyList[0] = keySearch
	}
	return keyList
}

func (r *BaseRepositoryImpl[T]) FindUserByEmail(
	ctx context.Context, tx *gorm.DB, email string,
) (*T, error) {
	var data T
	if err := tx.WithContext(ctx).Preload(clause.Associations).Where("email = ?", email).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		slog.Error("failed to find by email", slog.Any("error", err))
		return nil, err
	}
	return &data, nil
}
