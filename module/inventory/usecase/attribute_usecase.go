package usecase

// import (
// 	"context"
// 	"errors"

// 	"gitlab.com/integra_sm/cherry-v2-core-service/model"
// 	"gitlab.com/integra_sm/cherry-v2-core-service/module/attribute/repository/postgres"

// 	"gorm.io/gorm"
// )

// type AttributeValueUseCaseImpl struct {
// 	db                 *gorm.DB
// 	repoAttributeValue postgres.AttributeValueRepository
// 	repoAttribute      postgres.AttributeRepository
// }

// func NewAttributeUseCase(
// 	db *gorm.DB, repoAttributeValue postgres.AttributeValueRepository, repoAttribute postgres.AttributeRepository,
// ) AttributeValueUseCase {
// 	return &AttributeValueUseCaseImpl{
// 		db:                 db,
// 		repoAttributeValue: repoAttributeValue,
// 		repoAttribute:      repoAttribute,
// 	}
// }

// func (s *AttributeValueUseCaseImpl) Create(
// 	ctx context.Context, req *model.CreateAttributeValueReq,
// ) (*model.CreateAttributeValueRes, error) {
// 	tx := s.db.Begin()
// 	defer tx.Rollback()

// 	body := req.ToDomain(ctx)

// 	if err := s.repoAttributeValue.CreateTx(ctx, tx, body); err != nil {
// 		return nil, err
// 	}

// 	if err := tx.Commit().Error; err != nil {
// 		return nil, err
// 	}
// 	return &model.CreateAttributeValueRes{AttributeValue: body}, nil
// }

// func (s *AttributeValueUseCaseImpl) GetById(
// 	ctx context.Context, req *model.GetIdAttributeValueReq,
// ) (*model.GetIdAttributeValueRes, error) {
// 	result, err := s.repoAttributeValue.FindAttributeValue(ctx, s.db, req.ReferencesId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if result == nil {
// 		return nil, errors.New("attribute not found")
// 	}

// 	return &model.GetIdAttributeValueRes{AttributeValue: result}, nil
// }

// func (s *AttributeValueUseCaseImpl) Update(
// 	ctx context.Context, req *model.UpdateAttributeValueReq,
// ) (*model.UpdateAttributeValueRes, error) {
// 	tx := s.db.Begin()
// 	defer tx.Rollback()

// 	body := req.ToDomain(ctx)

// 	data, err := s.repoAttributeValue.FindByReferencesID(ctx, s.db, req.ReferencesId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if data == nil {
// 		return nil, errors.New("attribute not found")
// 	}

// 	body.Id = data.Id

// 	if err := s.repoAttributeValue.UpdateTx(ctx, tx, body); err != nil {
// 		return nil, err
// 	}

// 	if err := tx.Commit().Error; err != nil {
// 		return nil, err
// 	}
// 	return &model.UpdateAttributeValueRes{AttributeValue: body}, nil

// }

// func (s *AttributeValueUseCaseImpl) GetList(
// 	ctx context.Context, req *model.GetListAttributeValueReq,
// ) (*model.GetListAttributeValueRes, error) {
// 	result, err := s.repoAttributeValue.AttributeValueList(ctx, s.db, req.Page, req.Order, req.Filter, req.Keyword)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.GetListAttributeValueRes{
// 		Data:       result.Data,
// 		Pagination: result.ToPagination(),
// 	}, nil
// }

// func (s *AttributeValueUseCaseImpl) Delete(
// 	ctx context.Context, req *model.DeleteAttributeValueReq,
// ) (*model.DeleteAttributeValueRes, error) {
// 	tx := s.db.Begin()
// 	defer tx.Rollback()

// 	if err := s.repoAttributeValue.DeleteByReferencesIdTxBatch(ctx, tx, req.ReferencesId); err != nil {
// 		return nil, err
// 	}

// 	if err := tx.Commit().Error; err != nil {
// 		return nil, err
// 	}
// 	return &model.DeleteAttributeValueRes{ReferencesId: req.ReferencesId}, nil
// }

// func (s *AttributeValueUseCaseImpl) GetAttributePersonalData(
// 	ctx context.Context, req *model.GetListAttributeReq,
// ) (*model.GetListAttributeRes, error) {
// 	req.Filter = append(req.Filter, &model.FilterParam{
// 		Field:    "sub_module",
// 		Value:    "personal_data",
// 		Operator: "=",
// 	})
// 	result, err := s.repoAttribute.AttributeList(ctx, s.db, req.Page, req.Order, req.Filter, req.Keyword)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.GetListAttributeRes{
// 		Data:       result.Data,
// 		Pagination: result.ToPagination(),
// 	}, nil
// }

// func (s *AttributeValueUseCaseImpl) GetAttributeAddress(
// 	ctx context.Context, req *model.GetListAttributeReq,
// ) (*model.GetListAttributeRes, error) {
// 	req.Filter = append(req.Filter, &model.FilterParam{
// 		Field:    "sub_module",
// 		Value:    "address",
// 		Operator: "=",
// 	})
// 	result, err := s.repoAttribute.AttributeList(ctx, s.db, req.Page, req.Order, req.Filter, req.Keyword)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.GetListAttributeRes{
// 		Data:       result.Data,
// 		Pagination: result.ToPagination(),
// 	}, nil
// }

// func (s *AttributeValueUseCaseImpl) GetAttributeRelative(
// 	ctx context.Context, req *model.GetListAttributeReq,
// ) (*model.GetListAttributeRes, error) {
// 	req.Filter = append(req.Filter, &model.FilterParam{
// 		Field:    "sub_module",
// 		Value:    "relative",
// 		Operator: "=",
// 	})
// 	result, err := s.repoAttribute.AttributeList(ctx, s.db, req.Page, req.Order, req.Filter, req.Keyword)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.GetListAttributeRes{
// 		Data:       result.Data,
// 		Pagination: result.ToPagination(),
// 	}, nil
// }

// func (s *AttributeValueUseCaseImpl) GetAttributeEducation(
// 	ctx context.Context, req *model.GetListAttributeReq,
// ) (*model.GetListAttributeRes, error) {
// 	req.Filter = append(req.Filter, &model.FilterParam{
// 		Field:    "sub_module",
// 		Value:    "education",
// 		Operator: "=",
// 	})
// 	result, err := s.repoAttribute.AttributeList(ctx, s.db, req.Page, req.Order, req.Filter, req.Keyword)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.GetListAttributeRes{
// 		Data:       result.Data,
// 		Pagination: result.ToPagination(),
// 	}, nil
// }
