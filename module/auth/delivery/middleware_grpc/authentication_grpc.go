package middlewaregrpc

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"inventory-service/helpers/constant"
	"inventory-service/module/auth/usecase"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthenticationJWT struct {
	authUsecase usecase.AuthUsecase
	method      map[string][]string
}

func NewAuthenticationJWT(authUsecase usecase.AuthUsecase, method map[string][]string) AuthenticationJWT {
	return AuthenticationJWT{
		authUsecase: authUsecase,
		method:      method,
	}
}

func (n *AuthenticationJWT) JwtInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	method := strings.Split(info.FullMethod, "/")
	if len(method) < 3 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid method format: %s", info.FullMethod)
	}

	service := method[1]
	methodName := method[2]

	// Cek apakah methodName termasuk dalam whitelist
	if isWhitelisted(service, methodName, n.method) {
		return handler(ctx, req)
	}

	// Else: require JWT
	var err error
	ctx, err = n.SetMetaDataContext(ctx)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func isWhitelisted(service, method string, whitelist map[string][]string) bool {
	allowedMethods, ok := whitelist[service]
	if !ok {
		return false
	}
	for _, m := range allowedMethods {
		if m == method {
			return true
		}
	}
	return false
}

func (n *AuthenticationJWT) SetMetaDataContext(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	// Get the token from the metadata (usually under "authorization" key)
	tokenString := n.getTokenFromMetadata(md)

	if tokenString == "" {
		return nil, fmt.Errorf("missing token")
	}

	verifyResult, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !verifyResult.Valid {
		return nil, errors.New("invalid token")
	}

	claims := verifyResult.Claims.(jwt.MapClaims)

	// idUser := model.AuthCheckExist{Id: claims["user_id"].(string)}

	// ctx = constant.SetUser(ctx, global.StructToJson(user))

	ctx = n.setJWTMapClaims(claims, ctx)
	ctx = constant.SetToken(ctx, tokenString)

	return ctx, nil
}

func (n *AuthenticationJWT) GetMetadataAsContext(ctx context.Context, md metadata.MD) (context.Context, error) {
	tokenString := n.getTokenFromMetadata(md)

	if tokenString == "" {
		return ctx, fmt.Errorf("missing token")
	}

	verifyResult, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_SECRET")), nil
	})

	if err != nil {
		return ctx, err
	}

	if !verifyResult.Valid {
		return ctx, errors.New("invalid token")
	}

	claims := verifyResult.Claims.(jwt.MapClaims)
	ctx = n.setJWTMapClaims(claims, ctx)
	ctx = constant.SetToken(ctx, tokenString)
	// Log the claims for demonstration

	// Proceed to handle the actual gRPC request
	// return handler(ctx, req)
	return ctx, nil
}

// Validates the JWT and returns the claims if valid
func (n *AuthenticationJWT) getTokenFromMetadata(md metadata.MD) string {
	if authHeader, ok := md["authorization"]; ok && len(authHeader) > 0 {
		// Typically, the JWT token is prefixed by "Bearer "
		return authHeader[0][7:]
	}
	return ""
}

func (n *AuthenticationJWT) setJWTMapClaims(claims jwt.MapClaims, ctx context.Context) context.Context {
	// if claims["companies_references_id"] == nil {
	// 	ctx = constant.SetCompaniesRefId(ctx, nil)
	// } else {
	// 	companiesRefId := model.CheckCompaniesExist{CompaniesReferencesId: claims["companies_references_id"].(string)}
	// 	ctx = constant.SetCompaniesRefId(ctx, &companiesRefId.CompaniesReferencesId)
	// }

	// if claims["employee_references_id"] != nil {
	// 	employeeRefId := claims["employee_references_id"].(string)
	// 	ctx = constant.SetUserEmployeeRefId(ctx, &employeeRefId)
	// }

	// if claims["people_references_id"] != nil {
	// 	peopleRefId := claims["people_references_id"].(string)
	// 	ctx = constant.SetPeopleRefId(ctx, &peopleRefId)
	// }
	return ctx
}
