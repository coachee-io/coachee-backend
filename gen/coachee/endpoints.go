// Code generated by goa v3.0.7, DO NOT EDIT.
//
// coachee endpoints
//
// Command:
// $ goa gen coachee-backend/design

package coachee

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "coachee" service endpoints.
type Endpoints struct {
	GetCoaches          goa.Endpoint
	GetCoach            goa.Endpoint
	LenCoaches          goa.Endpoint
	CreateCoach         goa.Endpoint
	UpdateCoach         goa.Endpoint
	CreateCertification goa.Endpoint
	DeleteCertification goa.Endpoint
	CreateProgram       goa.Endpoint
	DeleteProgram       goa.Endpoint
	CreateAvailability  goa.Endpoint
	DeleteAvailability  goa.Endpoint
}

// NewEndpoints wraps the methods of the "coachee" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetCoaches:          NewGetCoachesEndpoint(s),
		GetCoach:            NewGetCoachEndpoint(s),
		LenCoaches:          NewLenCoachesEndpoint(s),
		CreateCoach:         NewCreateCoachEndpoint(s),
		UpdateCoach:         NewUpdateCoachEndpoint(s),
		CreateCertification: NewCreateCertificationEndpoint(s),
		DeleteCertification: NewDeleteCertificationEndpoint(s),
		CreateProgram:       NewCreateProgramEndpoint(s),
		DeleteProgram:       NewDeleteProgramEndpoint(s),
		CreateAvailability:  NewCreateAvailabilityEndpoint(s),
		DeleteAvailability:  NewDeleteAvailabilityEndpoint(s),
	}
}

// Use applies the given middleware to all the "coachee" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetCoaches = m(e.GetCoaches)
	e.GetCoach = m(e.GetCoach)
	e.LenCoaches = m(e.LenCoaches)
	e.CreateCoach = m(e.CreateCoach)
	e.UpdateCoach = m(e.UpdateCoach)
	e.CreateCertification = m(e.CreateCertification)
	e.DeleteCertification = m(e.DeleteCertification)
	e.CreateProgram = m(e.CreateProgram)
	e.DeleteProgram = m(e.DeleteProgram)
	e.CreateAvailability = m(e.CreateAvailability)
	e.DeleteAvailability = m(e.DeleteAvailability)
}

// NewGetCoachesEndpoint returns an endpoint function that calls the method
// "GetCoaches" of service "coachee".
func NewGetCoachesEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetCoachesPayload)
		return s.GetCoaches(ctx, p)
	}
}

// NewGetCoachEndpoint returns an endpoint function that calls the method
// "GetCoach" of service "coachee".
func NewGetCoachEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetCoachPayload)
		return s.GetCoach(ctx, p)
	}
}

// NewLenCoachesEndpoint returns an endpoint function that calls the method
// "LenCoaches" of service "coachee".
func NewLenCoachesEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LenCoachesPayload)
		return s.LenCoaches(ctx, p)
	}
}

// NewCreateCoachEndpoint returns an endpoint function that calls the method
// "CreateCoach" of service "coachee".
func NewCreateCoachEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateCoachPayload)
		return s.CreateCoach(ctx, p)
	}
}

// NewUpdateCoachEndpoint returns an endpoint function that calls the method
// "UpdateCoach" of service "coachee".
func NewUpdateCoachEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateCoachPayload)
		return nil, s.UpdateCoach(ctx, p)
	}
}

// NewCreateCertificationEndpoint returns an endpoint function that calls the
// method "CreateCertification" of service "coachee".
func NewCreateCertificationEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateCertificationPayload)
		return nil, s.CreateCertification(ctx, p)
	}
}

// NewDeleteCertificationEndpoint returns an endpoint function that calls the
// method "DeleteCertification" of service "coachee".
func NewDeleteCertificationEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteCertificationPayload)
		return nil, s.DeleteCertification(ctx, p)
	}
}

// NewCreateProgramEndpoint returns an endpoint function that calls the method
// "CreateProgram" of service "coachee".
func NewCreateProgramEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateProgramPayload)
		return nil, s.CreateProgram(ctx, p)
	}
}

// NewDeleteProgramEndpoint returns an endpoint function that calls the method
// "DeleteProgram" of service "coachee".
func NewDeleteProgramEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteProgramPayload)
		return nil, s.DeleteProgram(ctx, p)
	}
}

// NewCreateAvailabilityEndpoint returns an endpoint function that calls the
// method "CreateAvailability" of service "coachee".
func NewCreateAvailabilityEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateAvailabilityPayload)
		return nil, s.CreateAvailability(ctx, p)
	}
}

// NewDeleteAvailabilityEndpoint returns an endpoint function that calls the
// method "DeleteAvailability" of service "coachee".
func NewDeleteAvailabilityEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteAvailabilityPayload)
		return nil, s.DeleteAvailability(ctx, p)
	}
}
