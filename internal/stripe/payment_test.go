package stripe_test

import (
	"coachee-backend/internal/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_CreateOrder(t *testing.T) {
	cli := StartClient(t)

	customer := &model.Customer{
		ID:        3,
		FirstName: "Eli",
		LastName:  "Viveiros",
		Email:     "eli@viveiros.com",
	}

	order := &model.Order{
		ID:         1,
		CoachID:    2,
		CustomerID: 3,
		ProgramID:  "program_id",
		Amount:     10000,
		TaxPercent: 2000,
		Status:     "created",
	}

	err := cli.CreateCustomer(customer)
	require.Nil(t, err)
	require.NotEmpty(t, customer.StripeID)

	secret, err := cli.CreatePaymentIntent(order, customer)
	require.Nil(t, err)
	require.NotEmpty(t, order.PaymentIntentID)
	require.NotEmpty(t, secret)
}
