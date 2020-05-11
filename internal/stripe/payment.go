package stripe

import (
	"coachee-backend/internal/model"
	"fmt"

	"github.com/stripe/stripe-go"
)

func (c *Client) CreatePaymentIntent(order *model.Order, customer *model.Customer, coachID string) (string, error) {
	l := c.logger.With().Str("service", "CreatePaymentIntent").Logger()
	l.Debug().Msg("stripe CreatePaymentIntent called")

	param := &stripe.PaymentIntentParams{
		Amount:              stripe.Int64(int64(order.Amount)),
		CaptureMethod:       stripe.String("manual"),
		Currency:            stripe.String("GBP"),
		Customer:            &customer.StripeID,
		OnBehalfOf:          &coachID,
		ReceiptEmail:        &customer.Email,
		StatementDescriptor: stripe.String("Coachee"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	}
	param.AddMetadata("orderID", fmt.Sprintf("%d", order.ID))
	param.AddMetadata("customerID", fmt.Sprintf("%d", order.CustomerID))
	param.AddMetadata("coachID", fmt.Sprintf("%d", order.CoachID))
	param.AddMetadata("programID", order.ProgramID)

	pi, err := c.stripe.PaymentIntents.New(param)
	if err != nil {
		return "", err
	}
	order.PaymentIntentID = pi.ID

	l.Debug().Msg("stripe CreatePaymentIntent finished")
	return pi.ClientSecret, nil
}
