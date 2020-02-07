package stripe

import (
	"github.com/stripe/stripe-go"
)

func (c *Client) RegisterStripeExpress(authCode string) (string, error) {
	l := c.logger.With().Str("service", "RegisterStripeExpress").Str("authCode", authCode).Logger()
	l.Debug().Msg("stripe RegisterStripeExpress called")

	params := &stripe.OAuthTokenParams{
		GrantType: stripe.String("authorization_code"),
		Code:      stripe.String(authCode),
	}

	token, err := c.stripe.OAuth.New(params)
	if err != nil {
		l.Error().Err(err).Msg("failed to register stripe express account")
		return "", err
	}

	return token.StripeUserID, nil
}

func (c *Client) LoginStripeExpress(stripeID string) (string, error) {
	l := c.logger.With().Str("service", "LoginStripeExpress").Logger()
	l.Debug().Msg("stripe RegisterStripeExpress called")

	params := &stripe.LoginLinkParams{
		Account: stripe.String(stripeID),
	}
	link, err := c.stripe.LoginLinks.New(params)
	if err != nil {
		l.Error().Err(err).Msg("failed to generate login link for stripe express")
		return "", err
	}

	return link.URL, nil
}
