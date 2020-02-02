package stripe_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_LoginStripeExpress(t *testing.T) {
	cli := StartClient(t)

	url, err := cli.LoginStripeExpress("acct_1G7NvNHmXOSD14ch")
	require.Nil(t, err)
	fmt.Println(url)
}
