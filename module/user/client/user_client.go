package client

import (
	"context"

	pb "inventory-service/protocgen/inventory/v1/core/user"

	"inventory-service/module/auth/client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type UserClients struct {
	User pb.UserServiceClient
}

type CredentialsGRPC struct {
	token string
}

func (c *CredentialsGRPC) CredentialClientInterceptor(
	ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
) error {

	// todo @adam bikin logger monitoring

	authCtx := metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+c.token)

	return invoker(authCtx, method, req, reply, cc, opts...)
}

func NewClientAuth(ctx context.Context, params client.AuthParams) (*UserClients, error) {
	credentials := &CredentialsGRPC{
		token: params.Token,
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(credentials.CredentialClientInterceptor),
	}

	grpcClient, err := grpc.NewClient(params.URL, opts...)
	if err != nil {
		return nil, err
	}

	return &UserClients{
		User: pb.NewUserServiceClient(grpcClient),
	}, nil
}
