package grpc

import (
	"context"
	"reflect"

	"inventory-service/model"
	baseGRPC "inventory-service/module/base/delivery/grpc"
	"inventory-service/module/orders/usecase"
	pb "inventory-service/protocgen/inventory/v1/core/orders"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrdersService struct {
	OrdersUseCase usecase.OrdersUseCase
	pb.UnimplementedOrderServiceServer
	baseGRPC.GRPCHandler
}

func NewOrdersService(grpcServer *grpc.Server, usecase usecase.OrdersUseCase) {
	attributeGrpc := &OrdersService{
		OrdersUseCase: usecase,
	}
	pb.RegisterOrderServiceServer(grpcServer, attributeGrpc)
}

func (srv *OrdersService) CreateOrders(
	ctx context.Context, req *pb.CreateOrdersRequest,
) (*pb.CreateOrdersResponse, error) {
	var (
		request  model.CreateOrdersReq
		response pb.CreateOrdersResponse
	)
	if err := srv.Transform(req.GetOrders(), &request.BaseOrders); err != nil {
		return nil, err
	}

	result, err := srv.OrdersUseCase.Create(ctx, &request)
	if err != nil {
		return nil, err
	}
	response.Meta = srv.ResponseOK("Orders successfully created")
	response.Orders = &pb.Order{}

	if err := srv.Transform(result.Orders, response.Orders); err != nil {
		return nil, err
	}

	return &response, nil
}

func (srv *OrdersService) UpdateOrders(
	ctx context.Context, req *pb.UpdateOrdersRequest,
) (*pb.UpdateOrdersResponse, error) {
	var (
		request  model.UpdateOrdersReq
		response pb.UpdateOrdersResponse
	)
	err := srv.Transform(req.GetOrders(), &request.BaseOrders)
	if err != nil {
		return nil, err
	}

	request.Id = req.Orders.Id
	result, err := srv.OrdersUseCase.Update(ctx, &request)

	if err != nil {
		return nil, err
	}
	response.Meta = srv.ResponseOK("Orders succesfully updated")
	response.Orders = &pb.Order{}
	if err := srv.Transform(result.Orders, &response.Orders); err != nil {
		return nil, err
	}

	return &response, nil

}

func (srv *OrdersService) GetListOrders(
	ctx context.Context, req *pb.ListOrderssRequest,
) (*pb.ListOrderssResponse, error) {
	var (
		list     pb.ListOrderssResponse
		request  model.GetListOrdersReq
		errParse error
	)
	request.Page, request.Order, request.Filter, request.Keyword, ctx, errParse = srv.ParseListParams(ctx, req.GetPagination().GetOffset(), req.GetPagination().GetLimit(), req.GetQuery().GetOrder(), req.GetQuery().GetFilter(), req.GetQuery().GetKeyword())
	if errParse != nil {
		return nil, status.Error(codes.PermissionDenied, errParse.Error())
	}
	findReq := &model.GetListOrdersReq{
		Page:    request.Page,
		Filter:  request.Filter,
		Order:   request.Order,
		Keyword: request.Keyword,
	}

	result, err := srv.OrdersUseCase.GetList(ctx, findReq)
	if err != nil {
		return nil, err
	}

	if err := srv.TransformSlice(reflect.ValueOf(result.Data), reflect.ValueOf(&list.Orderss).Elem()); err != nil {
		return nil, err
	}

	list.Meta = srv.ResponseOKPagination("Orders data retrieved")
	if err := srv.Transform(result.Pagination, list.Meta.Pagination); err != nil {
		return nil, srv.ResponseError(err)
	}

	return &list, nil
}

func (srv *OrdersService) GetOrdersById(
	ctx context.Context, req *pb.GetOrdersRequest,
) (*pb.GetOrdersResponse, error) {
	var (
		request  model.GetIdOrdersReq
		response pb.GetOrdersResponse
	)

	request.Id = req.GetId()
	result, err := srv.OrdersUseCase.GetById(ctx, &request)
	if err != nil {
		return nil, err
	}

	response.Meta = srv.ResponseOK("Orders data retrieved")
	response.Orders = &pb.Order{}
	if err := srv.Transform(result.Orders, response.Orders); err != nil {
		return nil, err
	}

	return &response, nil
}

func (srv *OrdersService) DeleteOrders(
	ctx context.Context, req *pb.DeleteOrdersRequest,
) (*pb.DeleteOrdersResponse, error) {
	var (
		response pb.DeleteOrdersResponse
	)

	result, err := srv.OrdersUseCase.Delete(ctx, &model.DeleteOrdersReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	response.Meta = srv.ResponseOK("Orders data succesfully deleted")
	response.Orders = &pb.Order{}
	if err := srv.Transform(result.Orders, response.Orders); err != nil {
		return nil, err
	}

	return &response, nil

}
