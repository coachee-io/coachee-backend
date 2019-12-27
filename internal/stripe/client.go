package stripe

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/stripe/stripe-go/client"
)

// Client is the stripe client
type Client struct {
	stripe client.API

	logger *zerolog.Logger
	appCtx context.Context
}

// NewClient initializes a stripe client
func NewClient(appCtx context.Context, key string) *Client {
	log := zerolog.Ctx(appCtx).With().Str("component", "stripeClient").Logger()
	cli := &Client{
		stripe: client.API{},
		logger: &log,
		appCtx: appCtx,
	}

	cli.stripe.Init(key, nil)
	return cli
}
