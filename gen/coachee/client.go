// Code generated by goa v3.0.7, DO NOT EDIT.
//
// coachee client
//
// Command:
// $ goa gen coachee-backend/design

package coachee

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "coachee" service client.
type Client struct {
	GetCoachesEndpoint          goa.Endpoint
	GetCoachEndpoint            goa.Endpoint
	LenCoachesEndpoint          goa.Endpoint
	CreateCoachEndpoint         goa.Endpoint
	UpdateCoachEndpoint         goa.Endpoint
	CreateCertificationEndpoint goa.Endpoint
	DeleteCertificationEndpoint goa.Endpoint
	CreateProgramEndpoint       goa.Endpoint
	DeleteProgramEndpoint       goa.Endpoint
	CreateAvailabilityEndpoint  goa.Endpoint
	DeleteAvailabilityEndpoint  goa.Endpoint
	CreateCustomerEndpoint      goa.Endpoint
	CustomerLoginEndpoint       goa.Endpoint
	CreateOrderEndpoint         goa.Endpoint
}

// NewClient initializes a "coachee" service client given the endpoints.
func NewClient(getCoaches, getCoach, lenCoaches, createCoach, updateCoach, createCertification, deleteCertification, createProgram, deleteProgram, createAvailability, deleteAvailability, createCustomer, customerLogin, createOrder goa.Endpoint) *Client {
	return &Client{
		GetCoachesEndpoint:          getCoaches,
		GetCoachEndpoint:            getCoach,
		LenCoachesEndpoint:          lenCoaches,
		CreateCoachEndpoint:         createCoach,
		UpdateCoachEndpoint:         updateCoach,
		CreateCertificationEndpoint: createCertification,
		DeleteCertificationEndpoint: deleteCertification,
		CreateProgramEndpoint:       createProgram,
		DeleteProgramEndpoint:       deleteProgram,
		CreateAvailabilityEndpoint:  createAvailability,
		DeleteAvailabilityEndpoint:  deleteAvailability,
		CreateCustomerEndpoint:      createCustomer,
		CustomerLoginEndpoint:       customerLogin,
		CreateOrderEndpoint:         createOrder,
	}
}

// GetCoaches calls the "GetCoaches" endpoint of the "coachee" service.
func (c *Client) GetCoaches(ctx context.Context, p *GetCoachesPayload) (res []*Coach, err error) {
	var ires interface{}
	ires, err = c.GetCoachesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*Coach), nil
}

// GetCoach calls the "GetCoach" endpoint of the "coachee" service.
func (c *Client) GetCoach(ctx context.Context, p *GetCoachPayload) (res *Coach, err error) {
	var ires interface{}
	ires, err = c.GetCoachEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Coach), nil
}

// LenCoaches calls the "LenCoaches" endpoint of the "coachee" service.
func (c *Client) LenCoaches(ctx context.Context, p *LenCoachesPayload) (res uint, err error) {
	var ires interface{}
	ires, err = c.LenCoachesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(uint), nil
}

// CreateCoach calls the "CreateCoach" endpoint of the "coachee" service.
func (c *Client) CreateCoach(ctx context.Context, p *CreateCoachPayload) (res uint, err error) {
	var ires interface{}
	ires, err = c.CreateCoachEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(uint), nil
}

// UpdateCoach calls the "UpdateCoach" endpoint of the "coachee" service.
func (c *Client) UpdateCoach(ctx context.Context, p *UpdateCoachPayload) (err error) {
	_, err = c.UpdateCoachEndpoint(ctx, p)
	return
}

// CreateCertification calls the "CreateCertification" endpoint of the
// "coachee" service.
func (c *Client) CreateCertification(ctx context.Context, p *CreateCertificationPayload) (err error) {
	_, err = c.CreateCertificationEndpoint(ctx, p)
	return
}

// DeleteCertification calls the "DeleteCertification" endpoint of the
// "coachee" service.
func (c *Client) DeleteCertification(ctx context.Context, p *DeleteCertificationPayload) (err error) {
	_, err = c.DeleteCertificationEndpoint(ctx, p)
	return
}

// CreateProgram calls the "CreateProgram" endpoint of the "coachee" service.
func (c *Client) CreateProgram(ctx context.Context, p *CreateProgramPayload) (err error) {
	_, err = c.CreateProgramEndpoint(ctx, p)
	return
}

// DeleteProgram calls the "DeleteProgram" endpoint of the "coachee" service.
func (c *Client) DeleteProgram(ctx context.Context, p *DeleteProgramPayload) (err error) {
	_, err = c.DeleteProgramEndpoint(ctx, p)
	return
}

// CreateAvailability calls the "CreateAvailability" endpoint of the "coachee"
// service.
func (c *Client) CreateAvailability(ctx context.Context, p *CreateAvailabilityPayload) (err error) {
	_, err = c.CreateAvailabilityEndpoint(ctx, p)
	return
}

// DeleteAvailability calls the "DeleteAvailability" endpoint of the "coachee"
// service.
func (c *Client) DeleteAvailability(ctx context.Context, p *DeleteAvailabilityPayload) (err error) {
	_, err = c.DeleteAvailabilityEndpoint(ctx, p)
	return
}

// CreateCustomer calls the "CreateCustomer" endpoint of the "coachee" service.
func (c *Client) CreateCustomer(ctx context.Context, p *CreateCustomerPayload) (res *CreateCustomerResult, err error) {
	var ires interface{}
	ires, err = c.CreateCustomerEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*CreateCustomerResult), nil
}

// CustomerLogin calls the "CustomerLogin" endpoint of the "coachee" service.
func (c *Client) CustomerLogin(ctx context.Context, p *CustomerLoginPayload) (res *CustomerLoginResult, err error) {
	var ires interface{}
	ires, err = c.CustomerLoginEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*CustomerLoginResult), nil
}

// CreateOrder calls the "CreateOrder" endpoint of the "coachee" service.
func (c *Client) CreateOrder(ctx context.Context, p *CreateOrderPayload) (res *CreateOrderResult, err error) {
	var ires interface{}
	ires, err = c.CreateOrderEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*CreateOrderResult), nil
}
