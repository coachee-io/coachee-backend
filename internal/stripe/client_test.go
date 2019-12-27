package stripe_test

import (
	"coachee-backend/internal/stripe"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func StartClient(t *testing.T) *stripe.Client {
	key := "sk_test_yKV7Mo9kSpokxpFvwxKRtbyd00knjXTpJh"

	cli := stripe.NewClient(context.Background(), key)
	require.NotNil(t, cli)

	return cli
}

func TestNewClient(t *testing.T) {
	_ = StartClient(t)
}
