package client

// import (
// 	"context"

// 	"gitlab.com/integra_sm/cherry-v2-core-service/module/auth/client"
// 	pb "gitlab.com/integra_sm/cherry-v2-core-service/protocgen/cherry/v1/core/attribute"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// 	"google.golang.org/grpc/metadata"
// )

// type AttributeClients struct {
// 	Attribute pb.AttributeServiceClient
// }

// type CredentialsGRPC struct {
// 	token string
// }

// func (c *CredentialsGRPC) CredentialClientInterceptor(
// 	ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn,
// 	invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
// ) error {

// 	// todo @adam bikin logger monitoring

// 	authCtx := metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+c.token)

// 	return invoker(authCtx, method, req, reply, cc, opts...)
// }

// func NewClientAttribute(ctx context.Context, params client.AuthParams) (*AttributeClients, error) {
// 	credentials := CredentialsGRPC{
// 		token: params.Token,
// 	}

// 	opts := []grpc.DialOption{
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 		grpc.WithChainUnaryInterceptor(credentials.CredentialClientInterceptor),
// 	}

// 	grpcClient, err := grpc.NewClient(params.URL, opts...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &AttributeClients{
// 		Attribute: pb.NewAttributeServiceClient(grpcClient),
// 	}, nil
// }
