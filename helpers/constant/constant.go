package constant

import "context"

type (
	accessTokenKey string
)

const (
	AccessTokenKey accessTokenKey = "AccessToken"
)

func SetToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, AccessTokenKey, token)
}
