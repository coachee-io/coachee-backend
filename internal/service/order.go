package service

import (
	"coachee-backend/gen/coachee"
	"context"
)

// CreateOrder creates a new order
func (s *Service) CreateOrder(context.Context, *coachee.CreateOrderPayload) (res *coachee.CreateOrderResult, err error) {
	panic("implement me")
}
