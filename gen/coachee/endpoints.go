// Code generated by goa v3.0.9, DO NOT EDIT.
//
// coachee endpoints
//
// Command:
// $ goa gen coachee-backend/design

package coachee

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "coachee" service endpoints.
type Endpoints struct {
	StripeWebhooks                    goa.Endpoint
	GetCoaches                        goa.Endpoint
	GetCoach                          goa.Endpoint
	AdminGetCoach                     goa.Endpoint
	LenCoaches                        goa.Endpoint
	CreateCoach                       goa.Endpoint
	LoginCoach                        goa.Endpoint
	StartCoachPasswordRecoveryFlow    goa.Endpoint
	CheckCoachPasswordRecoveryToken   goa.Endpoint
	FinalizeCoachPasswordRecoveryFlow goa.Endpoint
	UpdateCoach                       goa.Endpoint
	CreateCertification               goa.Endpoint
	DeleteCertification               goa.Endpoint
	CreateProgram                     goa.Endpoint
	DeleteProgram                     goa.Endpoint
	CreateAvailability                goa.Endpoint
	DeleteAvailability                goa.Endpoint
	CreateCustomer                    goa.Endpoint
	CustomerLogin                     goa.Endpoint
	StartPasswordRecoveryFlow         goa.Endpoint
	CheckPasswordRecoveryToken        goa.Endpoint
	FinalizePasswordRecoveryFlow      goa.Endpoint
	CreateOrder                       goa.Endpoint
	RegisterStripeExpress             goa.Endpoint
	AdminLogin                        goa.Endpoint
	RegisterNewsletterEmail           goa.Endpoint
}

// NewEndpoints wraps the methods of the "coachee" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		StripeWebhooks:                    NewStripeWebhooksEndpoint(s),
		GetCoaches:                        NewGetCoachesEndpoint(s),
		GetCoach:                          NewGetCoachEndpoint(s),
		AdminGetCoach:                     NewAdminGetCoachEndpoint(s, a.JWTAuth),
		LenCoaches:                        NewLenCoachesEndpoint(s),
		CreateCoach:                       NewCreateCoachEndpoint(s),
		LoginCoach:                        NewLoginCoachEndpoint(s),
		StartCoachPasswordRecoveryFlow:    NewStartCoachPasswordRecoveryFlowEndpoint(s),
		CheckCoachPasswordRecoveryToken:   NewCheckCoachPasswordRecoveryTokenEndpoint(s),
		FinalizeCoachPasswordRecoveryFlow: NewFinalizeCoachPasswordRecoveryFlowEndpoint(s),
		UpdateCoach:                       NewUpdateCoachEndpoint(s, a.JWTAuth),
		CreateCertification:               NewCreateCertificationEndpoint(s, a.JWTAuth),
		DeleteCertification:               NewDeleteCertificationEndpoint(s, a.JWTAuth),
		CreateProgram:                     NewCreateProgramEndpoint(s, a.JWTAuth),
		DeleteProgram:                     NewDeleteProgramEndpoint(s, a.JWTAuth),
		CreateAvailability:                NewCreateAvailabilityEndpoint(s, a.JWTAuth),
		DeleteAvailability:                NewDeleteAvailabilityEndpoint(s, a.JWTAuth),
		CreateCustomer:                    NewCreateCustomerEndpoint(s),
		CustomerLogin:                     NewCustomerLoginEndpoint(s),
		StartPasswordRecoveryFlow:         NewStartPasswordRecoveryFlowEndpoint(s),
		CheckPasswordRecoveryToken:        NewCheckPasswordRecoveryTokenEndpoint(s),
		FinalizePasswordRecoveryFlow:      NewFinalizePasswordRecoveryFlowEndpoint(s),
		CreateOrder:                       NewCreateOrderEndpoint(s, a.JWTAuth),
		RegisterStripeExpress:             NewRegisterStripeExpressEndpoint(s),
		AdminLogin:                        NewAdminLoginEndpoint(s),
		RegisterNewsletterEmail:           NewRegisterNewsletterEmailEndpoint(s),
	}
}

// Use applies the given middleware to all the "coachee" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.StripeWebhooks = m(e.StripeWebhooks)
	e.GetCoaches = m(e.GetCoaches)
	e.GetCoach = m(e.GetCoach)
	e.AdminGetCoach = m(e.AdminGetCoach)
	e.LenCoaches = m(e.LenCoaches)
	e.CreateCoach = m(e.CreateCoach)
	e.LoginCoach = m(e.LoginCoach)
	e.StartCoachPasswordRecoveryFlow = m(e.StartCoachPasswordRecoveryFlow)
	e.CheckCoachPasswordRecoveryToken = m(e.CheckCoachPasswordRecoveryToken)
	e.FinalizeCoachPasswordRecoveryFlow = m(e.FinalizeCoachPasswordRecoveryFlow)
	e.UpdateCoach = m(e.UpdateCoach)
	e.CreateCertification = m(e.CreateCertification)
	e.DeleteCertification = m(e.DeleteCertification)
	e.CreateProgram = m(e.CreateProgram)
	e.DeleteProgram = m(e.DeleteProgram)
	e.CreateAvailability = m(e.CreateAvailability)
	e.DeleteAvailability = m(e.DeleteAvailability)
	e.CreateCustomer = m(e.CreateCustomer)
	e.CustomerLogin = m(e.CustomerLogin)
	e.StartPasswordRecoveryFlow = m(e.StartPasswordRecoveryFlow)
	e.CheckPasswordRecoveryToken = m(e.CheckPasswordRecoveryToken)
	e.FinalizePasswordRecoveryFlow = m(e.FinalizePasswordRecoveryFlow)
	e.CreateOrder = m(e.CreateOrder)
	e.RegisterStripeExpress = m(e.RegisterStripeExpress)
	e.AdminLogin = m(e.AdminLogin)
	e.RegisterNewsletterEmail = m(e.RegisterNewsletterEmail)
}

// NewStripeWebhooksEndpoint returns an endpoint function that calls the method
// "StripeWebhooks" of service "coachee".
func NewStripeWebhooksEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(map[string]interface{})
		return nil, s.StripeWebhooks(ctx, p)
	}
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

// NewAdminGetCoachEndpoint returns an endpoint function that calls the method
// "AdminGetCoach" of service "coachee".
func NewAdminGetCoachEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AdminGetCoachPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"client", "admin"},
			RequiredScopes: []string{"admin"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.AdminGetCoach(ctx, p)
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

// NewLoginCoachEndpoint returns an endpoint function that calls the method
// "LoginCoach" of service "coachee".
func NewLoginCoachEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LoginCoachPayload)
		return s.LoginCoach(ctx, p)
	}
}

// NewStartCoachPasswordRecoveryFlowEndpoint returns an endpoint function that
// calls the method "StartCoachPasswordRecoveryFlow" of service "coachee".
func NewStartCoachPasswordRecoveryFlowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*StartCoachPasswordRecoveryFlowPayload)
		return nil, s.StartCoachPasswordRecoveryFlow(ctx, p)
	}
}

// NewCheckCoachPasswordRecoveryTokenEndpoint returns an endpoint function that
// calls the method "CheckCoachPasswordRecoveryToken" of service "coachee".
func NewCheckCoachPasswordRecoveryTokenEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CheckCoachPasswordRecoveryTokenPayload)
		return nil, s.CheckCoachPasswordRecoveryToken(ctx, p)
	}
}

// NewFinalizeCoachPasswordRecoveryFlowEndpoint returns an endpoint function
// that calls the method "FinalizeCoachPasswordRecoveryFlow" of service
// "coachee".
func NewFinalizeCoachPasswordRecoveryFlowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*FinalizeCoachPasswordRecoveryFlowPayload)
		return nil, s.FinalizeCoachPasswordRecoveryFlow(ctx, p)
	}
}

// NewUpdateCoachEndpoint returns an endpoint function that calls the method
// "UpdateCoach" of service "coachee".
func NewUpdateCoachEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateCoachPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"client", "admin"},
			RequiredScopes: []string{"admin"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.UpdateCoach(ctx, p)
	}
}

// NewCreateCertificationEndpoint returns an endpoint function that calls the
// method "CreateCertification" of service "coachee".
func NewCreateCertificationEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateCertificationPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"client", "admin"},
			RequiredScopes: []string{"admin"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.CreateCertification(ctx, p)
	}
}

// NewDeleteCertificationEndpoint returns an endpoint function that calls the
// method "DeleteCertification" of service "coachee".
func NewDeleteCertificationEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteCertificationPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"client", "admin"},
			RequiredScopes: []string{"admin"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.DeleteCertification(ctx, p)
	}
}

// NewCreateProgramEndpoint returns an endpoint function that calls the method
// "CreateProgram" of service "coachee".
func NewCreateProgramEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateProgramPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"client", "admin"},
			RequiredScopes: []string{"admin"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.CreateProgram(ctx, p)
	}
}

// NewDeleteProgramEndpoint returns an endpoint function that calls the method
// "DeleteProgram" of service "coachee".
func NewDeleteProgramEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteProgramPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"client", "admin"},
			RequiredScopes: []string{"admin"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.DeleteProgram(ctx, p)
	}
}

// NewCreateAvailabilityEndpoint returns an endpoint function that calls the
// method "CreateAvailability" of service "coachee".
func NewCreateAvailabilityEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateAvailabilityPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"client", "admin"},
			RequiredScopes: []string{"admin"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.CreateAvailability(ctx, p)
	}
}

// NewDeleteAvailabilityEndpoint returns an endpoint function that calls the
// method "DeleteAvailability" of service "coachee".
func NewDeleteAvailabilityEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteAvailabilityPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"client", "admin"},
			RequiredScopes: []string{"admin"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.DeleteAvailability(ctx, p)
	}
}

// NewCreateCustomerEndpoint returns an endpoint function that calls the method
// "CreateCustomer" of service "coachee".
func NewCreateCustomerEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateCustomerPayload)
		return s.CreateCustomer(ctx, p)
	}
}

// NewCustomerLoginEndpoint returns an endpoint function that calls the method
// "CustomerLogin" of service "coachee".
func NewCustomerLoginEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CustomerLoginPayload)
		return s.CustomerLogin(ctx, p)
	}
}

// NewStartPasswordRecoveryFlowEndpoint returns an endpoint function that calls
// the method "StartPasswordRecoveryFlow" of service "coachee".
func NewStartPasswordRecoveryFlowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*StartPasswordRecoveryFlowPayload)
		return nil, s.StartPasswordRecoveryFlow(ctx, p)
	}
}

// NewCheckPasswordRecoveryTokenEndpoint returns an endpoint function that
// calls the method "CheckPasswordRecoveryToken" of service "coachee".
func NewCheckPasswordRecoveryTokenEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CheckPasswordRecoveryTokenPayload)
		return nil, s.CheckPasswordRecoveryToken(ctx, p)
	}
}

// NewFinalizePasswordRecoveryFlowEndpoint returns an endpoint function that
// calls the method "FinalizePasswordRecoveryFlow" of service "coachee".
func NewFinalizePasswordRecoveryFlowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*FinalizePasswordRecoveryFlowPayload)
		return nil, s.FinalizePasswordRecoveryFlow(ctx, p)
	}
}

// NewCreateOrderEndpoint returns an endpoint function that calls the method
// "CreateOrder" of service "coachee".
func NewCreateOrderEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateOrderPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"client", "admin"},
			RequiredScopes: []string{"client"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.CreateOrder(ctx, p)
	}
}

// NewRegisterStripeExpressEndpoint returns an endpoint function that calls the
// method "RegisterStripeExpress" of service "coachee".
func NewRegisterStripeExpressEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RegisterStripeExpressPayload)
		return nil, s.RegisterStripeExpress(ctx, p)
	}
}

// NewAdminLoginEndpoint returns an endpoint function that calls the method
// "AdminLogin" of service "coachee".
func NewAdminLoginEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AdminLoginPayload)
		return s.AdminLogin(ctx, p)
	}
}

// NewRegisterNewsletterEmailEndpoint returns an endpoint function that calls
// the method "RegisterNewsletterEmail" of service "coachee".
func NewRegisterNewsletterEmailEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RegisterNewsletterEmailPayload)
		return nil, s.RegisterNewsletterEmail(ctx, p)
	}
}
