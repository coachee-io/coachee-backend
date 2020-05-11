package service

import (
	"coachee-backend/gen/coachee"
	"coachee-backend/internal/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/stripe/stripe-go"
)

// StripeWebhooks receives stripe webhooks
func (s *Service) StripeWebhooks(ctx context.Context, mapJSON map[string]interface{}) error {
	l := s.logger.With().Str("service", "StripeWebhooks").Logger()

	rawJSON, err := json.Marshal(mapJSON)
	if err != nil {
		l.Error().Err(err).Msg("failed to unmarshal mapJSON")
		return coachee.MakeInternal(err)
	}

	event := stripe.Event{}

	if err := json.Unmarshal(rawJSON, event); err != nil {
		l.Error().Err(err).Msg("failed to parse webhook body json")
		return coachee.MakeInternal(err)
	}

	switch event.Type {
	case "charge.succeeded":
		var charge stripe.Charge
		if err := json.Unmarshal(event.Data.Raw, &charge); err != nil {
			l.Error().Err(err).Msg("failed to parse charge.succeeded stripe event")
			return coachee.MakeInternal(err)
		}

		sOrderID := charge.Metadata["orderID"]
		id, err := strconv.ParseUint(sOrderID, 10, 64)
		if err != nil {
			l.Error().Err(err).Str("id", sOrderID).Msg("failed to retrieve sOrderID from charge.succeeded stripe event")
			return coachee.MakeInternal(err)
		}

		tx := s.orderRepository.Begin()
		order, err := s.orderRepository.GetByID(tx, uint(id))
		if err != nil {
			l.Error().Err(err).Uint64("orderID", id).Msg("failed to retrieve order")
			_ = tx.Rollback()
			return coachee.MakeInternal(err)
		}

		order.Status = model.OrderStatusCharged
		if err := s.orderRepository.Update(tx, order); err != nil {
			l.Error().Err(err).Uint64("orderID", id).Msg("failed to store order")
			_ = tx.Rollback()
			return coachee.MakeInternal(err)
		}

		customer, err := s.customerRepository.GetByID(tx, order.CustomerID)
		if err != nil {
			l.Error().Err(err).Uint("customerID", order.CustomerID).Msg("failed to retrieve customer")
			_ = tx.Rollback()
			return coachee.MakeInternal(err)
		}

		coach, err := s.coachRepository.GetByID(tx, order.CoachID)
		if err != nil {
			l.Error().Err(err).Uint("coachID", order.CoachID).Msg("failed to retrieve coach")
			_ = tx.Rollback()
			return coachee.MakeInternal(err)
		}

		coachName := fmt.Sprintf("%s %s", coach.FirstName, coach.LastName)

		var programme string
		for _, p := range coach.Programs {
			if p.ID == order.ProgramID {
				programme = p.Name
			}
		}

		if err := tx.Commit(); err != nil {
			l.Error().Err(err).Msg("failed to commit transaction")
			return coachee.MakeInternal(err)
		}

		if err := s.email.SendBookingEmail(customer.Email, programme, coachName); err != nil {
			l.Error().Err(err).Msg("failed to send booking email")
			return coachee.MakeInternal(err)
		}

		return nil
	default:
		l.Error().Str("event", event.Type).Msg("failed to send booking email")
		return coachee.MakeInternal(errors.New("failed to send booking email"))
	}
}
