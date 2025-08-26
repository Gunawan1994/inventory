package client

import (
	"context"

	"inventory-service/module/auth/client"
	pb "inventory-service/protocgen/inventory/v1/core/orders"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type OrdersClients struct {
	Inventory pb.OrderServiceClient
}

type CredentialsGRPC struct {
	token string
}

func (c *CredentialsGRPC) CredentialClientInterceptor(
	ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
) error {

	authCtx := metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+c.token)

	return invoker(authCtx, method, req, reply, cc, opts...)
}

func NewClientOrders(ctx context.Context, params client.AuthParams) (*OrdersClients, error) {
	credentials := CredentialsGRPC{
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

	return &OrdersClients{
		Inventory: pb.NewOrderServiceClient(grpcClient),
	}, nil
}
