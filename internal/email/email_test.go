package email_test

import (
	"coachee-backend/internal/email"
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func getTestEmailClient(t *testing.T) *email.Client {
	_ = os.Setenv("EMAIL_USERNAME", "joaotest76@gmail.com")
	_ = os.Setenv("EMAIL_PASSWORD", "Matematica123")
	_ = os.Setenv("EMAIL_PATH", "../../web/tmpl/")

	cli, err := email.NewClient(context.Background(), "localhost")
	require.Nil(t, err)

	return cli
}

func TestNewClient(t *testing.T) {
	_ = getTestEmailClient(t)
}
