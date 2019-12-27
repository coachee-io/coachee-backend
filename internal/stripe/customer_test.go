package stripe_test

import (
	"coachee-backend/internal/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_CreateCustomer(t *testing.T) {
	cli := StartClient(t)

	customer := &model.Customer{
		ID:        99999999,
		FirstName: "Eli",
		LastName:  "Viveiros",
		Email:     "eli@viveiros.com",
	}

	err := cli.CreateCustomer(customer)
	require.Nil(t, err)
	require.NotEmpty(t, customer.StripeID)
}
