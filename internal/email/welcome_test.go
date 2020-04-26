// +build email

package email_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_SendWelcomeEmail(t *testing.T) {
	cli := getTestEmailClient(t)

	err := cli.SendWelcomeEmail("joca14@gmail.com", "test_token")
	require.Nil(t, err)
}
