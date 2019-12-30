package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/auth"
	"context"
	"errors"

	"goa.design/goa/v3/security"
)

var (
	ErrInvalidToken error = coachee.MakeUnauthorized(errors.New("jwtauth: invalid token"))
)

func (s *Service) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	l := s.logger.With().Str("middleware", "JWTAuth").Logger()

	claims, err := auth.ParseToken(token)
	if err != nil {
		l.Info().Err(err).Msg("failed to parse token")
		return ctx, err
	}

	if err := claims.Valid(); err != nil {
		l.Info().Err(err).Msg("invalid token")
		return nil, coachee.MakeUnauthorized(err)
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		l.Info().Msg("scopes is nil")
		return ctx, ErrInvalidToken
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		l.Info().Msg("no scopes found")
		return ctx, ErrInvalidToken
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := schema.Validate(scopesInToken); err != nil {
		l.Info().Err(err).Msg("missing scope for access")
		return ctx, ErrInvalidToken
	}

	return ctx, nil
}
