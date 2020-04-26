// +build email

package email_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_SendClientPasswordRecoveryEmail(t *testing.T) {
	cli := getTestEmailClient(t)

	err := cli.SendClientPasswordRecoveryEmail("joca14@gmail.com", "test_password_token")
	require.Nil(t, err)
}

func TestClient_SendCoachPasswordRecoveryEmail(t *testing.T) {
	cli := getTestEmailClient(t)

	err := cli.SendCoachPasswordRecoveryEmail("joca14@gmail.com", "test_password_token")
	require.Nil(t, err)
}
