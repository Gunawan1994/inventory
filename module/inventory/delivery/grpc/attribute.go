package grpc

// import (
// 	"context"
// 	"reflect"

// 	"gitlab.com/integra_sm/cherry-v2-core-service/model"
// 	"gitlab.com/integra_sm/cherry-v2-core-service/module/attribute/usecase"
// 	baseGRPC "gitlab.com/integra_sm/cherry-v2-core-service/module/base/delivery/grpc"
// 	pb "gitlab.com/integra_sm/cherry-v2-core-service/protocgen/cherry/v1/core/attribute"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// type AttributeService struct {
// 	AttributeUseCase usecase.AttributeValueUseCase
// 	pb.UnimplementedAttributeServiceServer
// 	baseGRPC.GRPCHandler
// }

// func NewAttributeService(grpcServer *grpc.Server, usecase usecase.AttributeValueUseCase) {
// 	attributeGrpc := &AttributeService{
// 		AttributeUseCase: usecase,
// 	}
// 	pb.RegisterAttributeServiceServer(grpcServer, attributeGrpc)
// }

// func (srv *AttributeService) CreateAttributeValue(
// 	ctx context.Context, req *pb.CreateAttributeValueRequest,
// ) (*pb.CreateAttributeValueResponse, error) {
// 	var (
// 		request  model.CreateAttributeValueReq
// 		response pb.CreateAttributeValueResponse
// 	)
// 	if err := srv.Transform(req.GetValue(), &request.BaseAttributeValue); err != nil {
// 		return nil, err
// 	}

// 	result, err := srv.AttributeUseCase.Create(ctx, &request)

// 	if err != nil {
// 		return nil, err
// 	}
// 	response.Meta = srv.ResponseOK("Attribute successfully created")
// 	response.Value = &pb.AttributeValue{}

// 	if err := srv.Transform(result.AttributeValue, response.Value); err != nil {
// 		return nil, err
// 	}

// 	return &response, nil
// }

// func (srv *AttributeService) UpdateAttributeValue(
// 	ctx context.Context, req *pb.UpdateAttributeValueRequest,
// ) (*pb.UpdateAttributeValueResponse, error) {
// 	var (
// 		request  model.UpdateAttributeValueReq
// 		response pb.UpdateAttributeValueResponse
// 	)
// 	err := srv.Transform(req.GetValue(), &request.BaseAttributeValue)
// 	if err != nil {
// 		return nil, err
// 	}
// 	request.ReferencesId = req.ReferencesId
// 	result, err := srv.AttributeUseCase.Update(ctx, &request)

// 	if err != nil {
// 		return nil, err
// 	}
// 	response.Meta = srv.ResponseOK("Attribute succesfully updated")
// 	response.Value = &pb.AttributeValue{}
// 	if err := srv.Transform(result.AttributeValue, &response.Value); err != nil {
// 		return nil, err
// 	}

// 	return &response, nil

// }

// func (srv *AttributeService) GetListAttributeValue(
// 	ctx context.Context, req *pb.GetListAttributeValueRequest,
// ) (*pb.GetListAttributeValueResponse, error) {
// 	var (
// 		listAttribute pb.GetListAttributeValueResponse
// 		request       model.ListReq
// 		errParse      error
// 	)
// 	request.Page, request.Order, request.Filter, request.Keyword, ctx, errParse = srv.ParseListParams(ctx, req.GetPagination().GetOffset(), req.GetPagination().GetLimit(), req.GetQuery().GetOrder(), req.GetQuery().GetFilter(), req.GetQuery().GetKeyword())
// 	if errParse != nil {
// 		return nil, status.Error(codes.PermissionDenied, errParse.Error())
// 	}
// 	findReq := &model.GetListAttributeValueReq{
// 		Page:    request.Page,
// 		Filter:  request.Filter,
// 		Order:   request.Order,
// 		Keyword: request.Keyword,
// 	}
// 	result, err := srv.AttributeUseCase.GetList(ctx, findReq)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := srv.TransformSlice(reflect.ValueOf(result.Data), reflect.ValueOf(&listAttribute.Value).Elem()); err != nil {
// 		return nil, err
// 	}

// 	listAttribute.Meta = srv.ResponseOKPagination("Attribute data retrieved")
// 	if err := srv.Transform(result.Pagination, listAttribute.Meta.Pagination); err != nil {
// 		return nil, srv.ResponseError(err)
// 	}

// 	return &listAttribute, nil
// }

// func (srv *AttributeService) GetAttributeValueById(
// 	ctx context.Context, req *pb.AttributeValueGetByIdRequest,
// ) (*pb.AttributeValueGetByIdResponse, error) {
// 	var (
// 		getAttributeRequest model.GetIdAttributeValueReq
// 		getAttributeResult  pb.AttributeValueGetByIdResponse
// 	)

// 	getAttributeRequest.ReferencesId = req.GetReferencesId()
// 	result, err := srv.AttributeUseCase.GetById(ctx, &getAttributeRequest)

// 	if err != nil {
// 		return nil, err
// 	}

// 	getAttributeResult.Meta = srv.ResponseOK("Attribute data retrieved")
// 	getAttributeResult.Value = &pb.AttributeValue{}
// 	if err := srv.Transform(result.AttributeValue, getAttributeResult.Value); err != nil {
// 		return nil, err
// 	}
// 	return &getAttributeResult, nil
// }

// func (srv *AttributeService) DeleteAttributeValue(
// 	ctx context.Context, req *pb.DeleteAttributeValueRequest,
// ) (*pb.DeleteAttributeValueResponse, error) {
// 	var (
// 		response pb.DeleteAttributeValueResponse
// 	)

// 	result, err := srv.AttributeUseCase.Delete(ctx, &model.DeleteAttributeValueReq{
// 		ReferencesId: req.ReferencesId,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	response.Meta = srv.ResponseOK("Attribute succesfully deleted")
// 	response.ReferencesId = result.ReferencesId

// 	return &response, nil

// }

// func (srv *AttributeService) GetAttributePersonalData(
// 	ctx context.Context, req *pb.GetAttributePersonalDataRequest,
// ) (*pb.GetAttributePersonalDataResponse, error) {
// 	var (
// 		listAttribute pb.GetAttributePersonalDataResponse
// 		request       model.ListReq
// 		errParse      error
// 	)

// 	request.Page, request.Order, request.Filter, request.Keyword, ctx, errParse = srv.ParseListParams(ctx, req.GetPagination().GetOffset(), req.GetPagination().GetLimit(), req.GetQuery().GetOrder(), req.GetQuery().GetFilter(), req.GetQuery().GetKeyword())
// 	if errParse != nil {
// 		return nil, status.Error(codes.PermissionDenied, errParse.Error())
// 	}
// 	findReq := &model.GetListAttributeReq{
// 		Page:    request.Page,
// 		Filter:  request.Filter,
// 		Order:   request.Order,
// 		Keyword: request.Keyword,
// 	}

// 	result, err := srv.AttributeUseCase.GetAttributePersonalData(ctx, findReq)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := srv.TransformSlice(reflect.ValueOf(result.Data), reflect.ValueOf(&listAttribute.Attributes).Elem()); err != nil {
// 		return nil, err
// 	}

// 	listAttribute.Meta = srv.ResponseOKPagination("Attribute data retrieved")

// 	return &listAttribute, nil
// }

// func (srv *AttributeService) GetAttributeAddress(
// 	ctx context.Context, req *pb.GetAttributeAddressRequest,
// ) (*pb.GetAttributeAddressResponse, error) {
// 	var (
// 		listAttribute pb.GetAttributeAddressResponse
// 		request       model.ListReq
// 		errParse      error
// 	)

// 	request.Page, request.Order, request.Filter, request.Keyword, ctx, errParse = srv.ParseListParams(ctx, req.GetPagination().GetOffset(), req.GetPagination().GetLimit(), req.GetQuery().GetOrder(), req.GetQuery().GetFilter(), req.GetQuery().GetKeyword())
// 	if errParse != nil {
// 		return nil, status.Error(codes.PermissionDenied, errParse.Error())
// 	}
// 	findReq := &model.GetListAttributeReq{
// 		Page:    request.Page,
// 		Filter:  request.Filter,
// 		Order:   request.Order,
// 		Keyword: request.Keyword,
// 	}

// 	result, err := srv.AttributeUseCase.GetAttributeAddress(ctx, findReq)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := srv.TransformSlice(reflect.ValueOf(result.Data), reflect.ValueOf(&listAttribute.Attributes).Elem()); err != nil {
// 		return nil, err
// 	}

// 	listAttribute.Meta = srv.ResponseOKPagination("Attribute data retrieved")

// 	return &listAttribute, nil
// }

// func (srv *AttributeService) GetAttributeRelative(
// 	ctx context.Context, req *pb.GetAttributeRelativeRequest,
// ) (*pb.GetAttributeRelativeResponse, error) {
// 	var (
// 		listAttribute pb.GetAttributeRelativeResponse
// 		request       model.ListReq
// 		errParse      error
// 	)

// 	request.Page, request.Order, request.Filter, request.Keyword, ctx, errParse = srv.ParseListParams(ctx, req.GetPagination().GetOffset(), req.GetPagination().GetLimit(), req.GetQuery().GetOrder(), req.GetQuery().GetFilter(), req.GetQuery().GetKeyword())
// 	if errParse != nil {
// 		return nil, status.Error(codes.PermissionDenied, errParse.Error())
// 	}
// 	findReq := &model.GetListAttributeReq{
// 		Page:    request.Page,
// 		Filter:  request.Filter,
// 		Order:   request.Order,
// 		Keyword: request.Keyword,
// 	}

// 	result, err := srv.AttributeUseCase.GetAttributeRelative(ctx, findReq)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := srv.TransformSlice(reflect.ValueOf(result.Data), reflect.ValueOf(&listAttribute.Attributes).Elem()); err != nil {
// 		return nil, err
// 	}

// 	listAttribute.Meta = srv.ResponseOKPagination("Attribute data retrieved")

// 	return &listAttribute, nil
// }

// func (srv *AttributeService) GetAttributeEducation(
// 	ctx context.Context, req *pb.GetAttributeEducationRequest,
// ) (*pb.GetAttributeEducationResponse, error) {
// 	var (
// 		listAttribute pb.GetAttributeEducationResponse
// 		request       model.ListReq
// 		errParse      error
// 	)

// 	request.Page, request.Order, request.Filter, request.Keyword, ctx, errParse = srv.ParseListParams(ctx, req.GetPagination().GetOffset(), req.GetPagination().GetLimit(), req.GetQuery().GetOrder(), req.GetQuery().GetFilter(), req.GetQuery().GetKeyword())
// 	if errParse != nil {
// 		return nil, status.Error(codes.PermissionDenied, errParse.Error())
// 	}
// 	findReq := &model.GetListAttributeReq{
// 		Page:    request.Page,
// 		Filter:  request.Filter,
// 		Order:   request.Order,
// 		Keyword: request.Keyword,
// 	}

// 	result, err := srv.AttributeUseCase.GetAttributeEducation(ctx, findReq)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := srv.TransformSlice(reflect.ValueOf(result.Data), reflect.ValueOf(&listAttribute.Attributes).Elem()); err != nil {
// 		return nil, err
// 	}

// 	listAttribute.Meta = srv.ResponseOKPagination("Attribute data retrieved")

// 	return &listAttribute, nil
// }
