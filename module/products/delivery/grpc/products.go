package grpc

import (
	"context"
	"reflect"

	"inventory-service/model"
	baseGRPC "inventory-service/module/base/delivery/grpc"
	"inventory-service/module/products/usecase"
	pb "inventory-service/protocgen/inventory/v1/core/products"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductsService struct {
	ProductsUseCase usecase.ProductsUseCase
	pb.UnimplementedProductsServiceServer
	baseGRPC.GRPCHandler
}

func NewProductsService(grpcServer *grpc.Server, usecase usecase.ProductsUseCase) {
	attributeGrpc := &ProductsService{
		ProductsUseCase: usecase,
	}
	pb.RegisterProductsServiceServer(grpcServer, attributeGrpc)
}

func (srv *ProductsService) CreateProducts(
	ctx context.Context, req *pb.CreateProductsRequest,
) (*pb.CreateProductsResponse, error) {
	var (
		request  model.CreateProductsReq
		response pb.CreateProductsResponse
	)
	if err := srv.Transform(req.GetProducts(), &request.BaseProducts); err != nil {
		return nil, err
	}

	result, err := srv.ProductsUseCase.Create(ctx, &request)
	if err != nil {
		return nil, err
	}
	response.Meta = srv.ResponseOK("Products successfully created")
	response.Products = &pb.Products{}

	if err := srv.Transform(result.Products, response.Products); err != nil {
		return nil, err
	}

	return &response, nil
}

func (srv *ProductsService) UpdateProducts(
	ctx context.Context, req *pb.UpdateProductsRequest,
) (*pb.UpdateProductsResponse, error) {
	var (
		request  model.UpdateProductsReq
		response pb.UpdateProductsResponse
	)
	err := srv.Transform(req.GetProducts(), &request.BaseProducts)
	if err != nil {
		return nil, err
	}

	request.Id = req.Products.Id
	result, err := srv.ProductsUseCase.Update(ctx, &request)

	if err != nil {
		return nil, err
	}
	response.Meta = srv.ResponseOK("Products succesfully updated")
	response.Products = &pb.Products{}
	if err := srv.Transform(result.Products, &response.Products); err != nil {
		return nil, err
	}

	return &response, nil

}

func (srv *ProductsService) GetListProducts(
	ctx context.Context, req *pb.ListProductssRequest,
) (*pb.ListProductssResponse, error) {
	var (
		list     pb.ListProductssResponse
		request  model.GetListProductsReq
		errParse error
	)
	request.Page, request.Order, request.Filter, request.Keyword, ctx, errParse = srv.ParseListParams(ctx, req.GetPagination().GetOffset(), req.GetPagination().GetLimit(), req.GetQuery().GetOrder(), req.GetQuery().GetFilter(), req.GetQuery().GetKeyword())
	if errParse != nil {
		return nil, status.Error(codes.PermissionDenied, errParse.Error())
	}
	findReq := &model.GetListProductsReq{
		Page:    request.Page,
		Filter:  request.Filter,
		Order:   request.Order,
		Keyword: request.Keyword,
	}

	result, err := srv.ProductsUseCase.GetList(ctx, findReq)
	if err != nil {
		return nil, err
	}

	if err := srv.TransformSlice(reflect.ValueOf(result.Data), reflect.ValueOf(&list.Productss).Elem()); err != nil {
		return nil, err
	}

	list.Meta = srv.ResponseOKPagination("Products data retrieved")
	if err := srv.Transform(result.Pagination, list.Meta.Pagination); err != nil {
		return nil, srv.ResponseError(err)
	}

	return &list, nil
}

func (srv *ProductsService) GetProductsById(
	ctx context.Context, req *pb.GetProductsRequest,
) (*pb.GetProductsResponse, error) {
	var (
		request  model.GetIdProductsReq
		response pb.GetProductsResponse
	)

	request.Id = req.GetId()
	result, err := srv.ProductsUseCase.GetById(ctx, &request)
	if err != nil {
		return nil, err
	}

	response.Meta = srv.ResponseOK("Products data retrieved")
	response.Products = &pb.Products{}
	if err := srv.Transform(result.Products, response.Products); err != nil {
		return nil, err
	}

	return &response, nil
}

func (srv *ProductsService) DeleteProducts(
	ctx context.Context, req *pb.DeleteProductsRequest,
) (*pb.DeleteProductsResponse, error) {
	var (
		response pb.DeleteProductsResponse
	)

	result, err := srv.ProductsUseCase.Delete(ctx, &model.DeleteProductsReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	response.Meta = srv.ResponseOK("Products data succesfully deleted")
	response.Products = &pb.Products{}
	if err := srv.Transform(result.Products, response.Products); err != nil {
		return nil, err
	}

	return &response, nil

}
