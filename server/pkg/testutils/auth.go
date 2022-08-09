package testutils

import (
	"context"
)

func GetAuthenticatedContext(ctx context.Context, userId string) context.Context {
	newCtx := context.WithValue(ctx, "userId", userId)
	return newCtx
}
