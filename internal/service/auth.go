package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/auth"
	"context"
	"errors"

	"goa.design/goa/v3/security"
)

var (
	ErrInvalidTokenScopes error = coachee.MakeUnauthorized(errors.New("invalid token"))
)

func (s *Service) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	claims, err := auth.ParseToken(token)
	if err != nil {
		return ctx, err
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, ErrInvalidTokenScopes
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := schema.Validate(scopesInToken); err != nil {
		return ctx, ErrInvalidTokenScopes
	}
	return ctx, nil
}
