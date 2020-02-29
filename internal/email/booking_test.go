package email_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_SendBookingEmail(t *testing.T) {
	cli := getTestEmailClient(t)

	err := cli.SendBookingEmail("joca14@gmail.com", "mega programme", "super coach")
	require.Nil(t, err)
}
