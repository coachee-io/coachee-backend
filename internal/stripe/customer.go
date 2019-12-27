package stripe

import (
	"coachee-backend/internal/model"

	"fmt"

	"github.com/stripe/stripe-go"
)

// CreateCustomer creates a stripe customer
func (c *Client) CreateCustomer(customer *model.Customer) error {
	l := c.logger.With().Str("service", "CreateCustomer").Logger()
	l.Debug().Msg("stripe CreateCustomer called")

	name := customer.FirstName + " " + customer.LastName
	params := &stripe.CustomerParams{
		Email: &customer.Email,
		Name:  &name,
	}
	params.AddMetadata("id", fmt.Sprintf("%d", customer.ID))

	cus, err := c.stripe.Customers.New(params)
	if err != nil {
		return err
	}
	customer.StripeID = cus.ID

	l.Debug().Msg("stripe CreateCustomer finished")
	return nil
}
