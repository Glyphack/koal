package authapi

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func parseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("jwt_secret")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.StandardClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, err
}

func userClaimFromToken(claims jwt.MapClaims) string {
	return fmt.Sprint(claims["sub"])
}

func AuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}
	newCtx := context.WithValue(ctx, "userId", tokenInfo.Subject)

	return newCtx, nil
}
