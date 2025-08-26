package grpc

import (
	"context"
	"inventory-service/module/auth/usecase"
	baseGRPC "inventory-service/module/base/delivery/grpc"
	pb "inventory-service/protocgen/inventory/v1/core/auth"

	"inventory-service/model"

	"google.golang.org/grpc"
)

type AuthService struct {
	authUsecase usecase.AuthUsecase
	pb.UnimplementedAuthServiceServer
	baseGRPC.GRPCHandler
}

func NewAuthService(grpcServer *grpc.Server, usecase usecase.AuthUsecase) {
	authGrpc := &AuthService{authUsecase: usecase}
	pb.RegisterAuthServiceServer(grpcServer, authGrpc)
}

// func (srv *AuthService) RegisterUser(
// 	ctx context.Context, req *pb.RegisterRequest,
// ) (*pb.UserResponse, error) {
// 	var (
// 		response pb.UserResponse
// 	)

// 	result, err := srv.authUsecase.RegisterUser(ctx, &model.CreateUserReq{
// 		BaseUser: &model.BaseUser{
// 			Email:    req.GetEmail(),
// 			Username: req.GetUsername(),
// 			Password: req.GetPassword(),
// 		},
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := srv.Transform(result, &response); err != nil {
// 		return nil, err
// 	}

// 	return &response, nil
// }

func (srv *AuthService) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	var (
		checkVerifyToken model.VerifyCredential
		checkResponse    pb.LoginResponse
	)

	if err := srv.GRPCHandler.Transform(req, &checkVerifyToken); err != nil {
		return nil, err
	}

	result, err := srv.authUsecase.VerifyCredential(ctx, checkVerifyToken)

	if err != nil {
		return nil, err
	}

	if err := srv.GRPCHandler.Transform(result, &checkResponse); err != nil {
		return nil, err
	}

	return &checkResponse, nil
}
