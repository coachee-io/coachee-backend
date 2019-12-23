package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/auth"
	"context"
	"errors"
	"time"

	"goa.design/goa/v3/security"
)

var (
	ErrInvalidToken error = coachee.MakeUnauthorized(errors.New("invalid token"))

	ErrExpiredToken error = coachee.MakeUnauthorized(errors.New("token is expired"))
)

func (s *Service) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	claims, err := auth.ParseToken(token)
	if err != nil {
		return ctx, err
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		return ctx, ErrInvalidToken
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, ErrInvalidToken
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := schema.Validate(scopesInToken); err != nil {
		return ctx, ErrInvalidToken
	}

	// validate expiry token
	expiry, ok := claims["expiry"]
	if !ok {
		return ctx, ErrInvalidToken
	}
	exp, ok := expiry.(time.Time)
	if !ok {
		return ctx, ErrInvalidToken
	}
	if exp.After(time.Now()) {
		return ctx, ErrExpiredToken
	}

	return ctx, nil
}
