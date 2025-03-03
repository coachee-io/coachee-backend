// Code generated by goa v3.0.9, DO NOT EDIT.
//
// coachee client HTTP transport
//
// Command:
// $ goa gen coachee-backend/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the coachee service endpoint HTTP clients.
type Client struct {
	// StripeWebhooks Doer is the HTTP client used to make requests to the
	// StripeWebhooks endpoint.
	StripeWebhooksDoer goahttp.Doer

	// GetCoaches Doer is the HTTP client used to make requests to the GetCoaches
	// endpoint.
	GetCoachesDoer goahttp.Doer

	// GetCoach Doer is the HTTP client used to make requests to the GetCoach
	// endpoint.
	GetCoachDoer goahttp.Doer

	// AdminGetCoach Doer is the HTTP client used to make requests to the
	// AdminGetCoach endpoint.
	AdminGetCoachDoer goahttp.Doer

	// LenCoaches Doer is the HTTP client used to make requests to the LenCoaches
	// endpoint.
	LenCoachesDoer goahttp.Doer

	// CreateCoach Doer is the HTTP client used to make requests to the CreateCoach
	// endpoint.
	CreateCoachDoer goahttp.Doer

	// LoginCoach Doer is the HTTP client used to make requests to the LoginCoach
	// endpoint.
	LoginCoachDoer goahttp.Doer

	// StartCoachPasswordRecoveryFlow Doer is the HTTP client used to make requests
	// to the StartCoachPasswordRecoveryFlow endpoint.
	StartCoachPasswordRecoveryFlowDoer goahttp.Doer

	// CheckCoachPasswordRecoveryToken Doer is the HTTP client used to make
	// requests to the CheckCoachPasswordRecoveryToken endpoint.
	CheckCoachPasswordRecoveryTokenDoer goahttp.Doer

	// FinalizeCoachPasswordRecoveryFlow Doer is the HTTP client used to make
	// requests to the FinalizeCoachPasswordRecoveryFlow endpoint.
	FinalizeCoachPasswordRecoveryFlowDoer goahttp.Doer

	// UpdateCoach Doer is the HTTP client used to make requests to the UpdateCoach
	// endpoint.
	UpdateCoachDoer goahttp.Doer

	// CreateCertification Doer is the HTTP client used to make requests to the
	// CreateCertification endpoint.
	CreateCertificationDoer goahttp.Doer

	// DeleteCertification Doer is the HTTP client used to make requests to the
	// DeleteCertification endpoint.
	DeleteCertificationDoer goahttp.Doer

	// CreateProgram Doer is the HTTP client used to make requests to the
	// CreateProgram endpoint.
	CreateProgramDoer goahttp.Doer

	// DeleteProgram Doer is the HTTP client used to make requests to the
	// DeleteProgram endpoint.
	DeleteProgramDoer goahttp.Doer

	// CreateAvailability Doer is the HTTP client used to make requests to the
	// CreateAvailability endpoint.
	CreateAvailabilityDoer goahttp.Doer

	// DeleteAvailability Doer is the HTTP client used to make requests to the
	// DeleteAvailability endpoint.
	DeleteAvailabilityDoer goahttp.Doer

	// CreateCustomer Doer is the HTTP client used to make requests to the
	// CreateCustomer endpoint.
	CreateCustomerDoer goahttp.Doer

	// CustomerLogin Doer is the HTTP client used to make requests to the
	// CustomerLogin endpoint.
	CustomerLoginDoer goahttp.Doer

	// StartPasswordRecoveryFlow Doer is the HTTP client used to make requests to
	// the StartPasswordRecoveryFlow endpoint.
	StartPasswordRecoveryFlowDoer goahttp.Doer

	// CheckPasswordRecoveryToken Doer is the HTTP client used to make requests to
	// the CheckPasswordRecoveryToken endpoint.
	CheckPasswordRecoveryTokenDoer goahttp.Doer

	// FinalizePasswordRecoveryFlow Doer is the HTTP client used to make requests
	// to the FinalizePasswordRecoveryFlow endpoint.
	FinalizePasswordRecoveryFlowDoer goahttp.Doer

	// CreateOrder Doer is the HTTP client used to make requests to the CreateOrder
	// endpoint.
	CreateOrderDoer goahttp.Doer

	// RegisterStripeExpress Doer is the HTTP client used to make requests to the
	// RegisterStripeExpress endpoint.
	RegisterStripeExpressDoer goahttp.Doer

	// AdminLogin Doer is the HTTP client used to make requests to the AdminLogin
	// endpoint.
	AdminLoginDoer goahttp.Doer

	// RegisterNewsletterEmail Doer is the HTTP client used to make requests to the
	// RegisterNewsletterEmail endpoint.
	RegisterNewsletterEmailDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the coachee service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		StripeWebhooksDoer:                    doer,
		GetCoachesDoer:                        doer,
		GetCoachDoer:                          doer,
		AdminGetCoachDoer:                     doer,
		LenCoachesDoer:                        doer,
		CreateCoachDoer:                       doer,
		LoginCoachDoer:                        doer,
		StartCoachPasswordRecoveryFlowDoer:    doer,
		CheckCoachPasswordRecoveryTokenDoer:   doer,
		FinalizeCoachPasswordRecoveryFlowDoer: doer,
		UpdateCoachDoer:                       doer,
		CreateCertificationDoer:               doer,
		DeleteCertificationDoer:               doer,
		CreateProgramDoer:                     doer,
		DeleteProgramDoer:                     doer,
		CreateAvailabilityDoer:                doer,
		DeleteAvailabilityDoer:                doer,
		CreateCustomerDoer:                    doer,
		CustomerLoginDoer:                     doer,
		StartPasswordRecoveryFlowDoer:         doer,
		CheckPasswordRecoveryTokenDoer:        doer,
		FinalizePasswordRecoveryFlowDoer:      doer,
		CreateOrderDoer:                       doer,
		RegisterStripeExpressDoer:             doer,
		AdminLoginDoer:                        doer,
		RegisterNewsletterEmailDoer:           doer,
		RestoreResponseBody:                   restoreBody,
		scheme:                                scheme,
		host:                                  host,
		decoder:                               dec,
		encoder:                               enc,
	}
}

// StripeWebhooks returns an endpoint that makes HTTP requests to the coachee
// service StripeWebhooks server.
func (c *Client) StripeWebhooks() goa.Endpoint {
	var (
		encodeRequest  = EncodeStripeWebhooksRequest(c.encoder)
		decodeResponse = DecodeStripeWebhooksResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStripeWebhooksRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.StripeWebhooksDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "StripeWebhooks", err)
		}
		return decodeResponse(resp)
	}
}

// GetCoaches returns an endpoint that makes HTTP requests to the coachee
// service GetCoaches server.
func (c *Client) GetCoaches() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetCoachesRequest(c.encoder)
		decodeResponse = DecodeGetCoachesResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetCoachesRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetCoachesDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "GetCoaches", err)
		}
		return decodeResponse(resp)
	}
}

// GetCoach returns an endpoint that makes HTTP requests to the coachee service
// GetCoach server.
func (c *Client) GetCoach() goa.Endpoint {
	var (
		decodeResponse = DecodeGetCoachResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetCoachRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetCoachDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "GetCoach", err)
		}
		return decodeResponse(resp)
	}
}

// AdminGetCoach returns an endpoint that makes HTTP requests to the coachee
// service AdminGetCoach server.
func (c *Client) AdminGetCoach() goa.Endpoint {
	var (
		encodeRequest  = EncodeAdminGetCoachRequest(c.encoder)
		decodeResponse = DecodeAdminGetCoachResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAdminGetCoachRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AdminGetCoachDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "AdminGetCoach", err)
		}
		return decodeResponse(resp)
	}
}

// LenCoaches returns an endpoint that makes HTTP requests to the coachee
// service LenCoaches server.
func (c *Client) LenCoaches() goa.Endpoint {
	var (
		decodeResponse = DecodeLenCoachesResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildLenCoachesRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.LenCoachesDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "LenCoaches", err)
		}
		return decodeResponse(resp)
	}
}

// CreateCoach returns an endpoint that makes HTTP requests to the coachee
// service CreateCoach server.
func (c *Client) CreateCoach() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateCoachRequest(c.encoder)
		decodeResponse = DecodeCreateCoachResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateCoachRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateCoachDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "CreateCoach", err)
		}
		return decodeResponse(resp)
	}
}

// LoginCoach returns an endpoint that makes HTTP requests to the coachee
// service LoginCoach server.
func (c *Client) LoginCoach() goa.Endpoint {
	var (
		encodeRequest  = EncodeLoginCoachRequest(c.encoder)
		decodeResponse = DecodeLoginCoachResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildLoginCoachRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.LoginCoachDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "LoginCoach", err)
		}
		return decodeResponse(resp)
	}
}

// StartCoachPasswordRecoveryFlow returns an endpoint that makes HTTP requests
// to the coachee service StartCoachPasswordRecoveryFlow server.
func (c *Client) StartCoachPasswordRecoveryFlow() goa.Endpoint {
	var (
		encodeRequest  = EncodeStartCoachPasswordRecoveryFlowRequest(c.encoder)
		decodeResponse = DecodeStartCoachPasswordRecoveryFlowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStartCoachPasswordRecoveryFlowRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.StartCoachPasswordRecoveryFlowDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "StartCoachPasswordRecoveryFlow", err)
		}
		return decodeResponse(resp)
	}
}

// CheckCoachPasswordRecoveryToken returns an endpoint that makes HTTP requests
// to the coachee service CheckCoachPasswordRecoveryToken server.
func (c *Client) CheckCoachPasswordRecoveryToken() goa.Endpoint {
	var (
		decodeResponse = DecodeCheckCoachPasswordRecoveryTokenResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCheckCoachPasswordRecoveryTokenRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CheckCoachPasswordRecoveryTokenDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "CheckCoachPasswordRecoveryToken", err)
		}
		return decodeResponse(resp)
	}
}

// FinalizeCoachPasswordRecoveryFlow returns an endpoint that makes HTTP
// requests to the coachee service FinalizeCoachPasswordRecoveryFlow server.
func (c *Client) FinalizeCoachPasswordRecoveryFlow() goa.Endpoint {
	var (
		encodeRequest  = EncodeFinalizeCoachPasswordRecoveryFlowRequest(c.encoder)
		decodeResponse = DecodeFinalizeCoachPasswordRecoveryFlowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildFinalizeCoachPasswordRecoveryFlowRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.FinalizeCoachPasswordRecoveryFlowDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "FinalizeCoachPasswordRecoveryFlow", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateCoach returns an endpoint that makes HTTP requests to the coachee
// service UpdateCoach server.
func (c *Client) UpdateCoach() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateCoachRequest(c.encoder)
		decodeResponse = DecodeUpdateCoachResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateCoachRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateCoachDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "UpdateCoach", err)
		}
		return decodeResponse(resp)
	}
}

// CreateCertification returns an endpoint that makes HTTP requests to the
// coachee service CreateCertification server.
func (c *Client) CreateCertification() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateCertificationRequest(c.encoder)
		decodeResponse = DecodeCreateCertificationResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateCertificationRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateCertificationDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "CreateCertification", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteCertification returns an endpoint that makes HTTP requests to the
// coachee service DeleteCertification server.
func (c *Client) DeleteCertification() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteCertificationRequest(c.encoder)
		decodeResponse = DecodeDeleteCertificationResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteCertificationRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteCertificationDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "DeleteCertification", err)
		}
		return decodeResponse(resp)
	}
}

// CreateProgram returns an endpoint that makes HTTP requests to the coachee
// service CreateProgram server.
func (c *Client) CreateProgram() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateProgramRequest(c.encoder)
		decodeResponse = DecodeCreateProgramResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateProgramRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateProgramDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "CreateProgram", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteProgram returns an endpoint that makes HTTP requests to the coachee
// service DeleteProgram server.
func (c *Client) DeleteProgram() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteProgramRequest(c.encoder)
		decodeResponse = DecodeDeleteProgramResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteProgramRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteProgramDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "DeleteProgram", err)
		}
		return decodeResponse(resp)
	}
}

// CreateAvailability returns an endpoint that makes HTTP requests to the
// coachee service CreateAvailability server.
func (c *Client) CreateAvailability() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateAvailabilityRequest(c.encoder)
		decodeResponse = DecodeCreateAvailabilityResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateAvailabilityRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateAvailabilityDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "CreateAvailability", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteAvailability returns an endpoint that makes HTTP requests to the
// coachee service DeleteAvailability server.
func (c *Client) DeleteAvailability() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteAvailabilityRequest(c.encoder)
		decodeResponse = DecodeDeleteAvailabilityResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteAvailabilityRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteAvailabilityDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "DeleteAvailability", err)
		}
		return decodeResponse(resp)
	}
}

// CreateCustomer returns an endpoint that makes HTTP requests to the coachee
// service CreateCustomer server.
func (c *Client) CreateCustomer() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateCustomerRequest(c.encoder)
		decodeResponse = DecodeCreateCustomerResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateCustomerRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateCustomerDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "CreateCustomer", err)
		}
		return decodeResponse(resp)
	}
}

// CustomerLogin returns an endpoint that makes HTTP requests to the coachee
// service CustomerLogin server.
func (c *Client) CustomerLogin() goa.Endpoint {
	var (
		encodeRequest  = EncodeCustomerLoginRequest(c.encoder)
		decodeResponse = DecodeCustomerLoginResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCustomerLoginRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CustomerLoginDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "CustomerLogin", err)
		}
		return decodeResponse(resp)
	}
}

// StartPasswordRecoveryFlow returns an endpoint that makes HTTP requests to
// the coachee service StartPasswordRecoveryFlow server.
func (c *Client) StartPasswordRecoveryFlow() goa.Endpoint {
	var (
		encodeRequest  = EncodeStartPasswordRecoveryFlowRequest(c.encoder)
		decodeResponse = DecodeStartPasswordRecoveryFlowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStartPasswordRecoveryFlowRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.StartPasswordRecoveryFlowDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "StartPasswordRecoveryFlow", err)
		}
		return decodeResponse(resp)
	}
}

// CheckPasswordRecoveryToken returns an endpoint that makes HTTP requests to
// the coachee service CheckPasswordRecoveryToken server.
func (c *Client) CheckPasswordRecoveryToken() goa.Endpoint {
	var (
		decodeResponse = DecodeCheckPasswordRecoveryTokenResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCheckPasswordRecoveryTokenRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CheckPasswordRecoveryTokenDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "CheckPasswordRecoveryToken", err)
		}
		return decodeResponse(resp)
	}
}

// FinalizePasswordRecoveryFlow returns an endpoint that makes HTTP requests to
// the coachee service FinalizePasswordRecoveryFlow server.
func (c *Client) FinalizePasswordRecoveryFlow() goa.Endpoint {
	var (
		encodeRequest  = EncodeFinalizePasswordRecoveryFlowRequest(c.encoder)
		decodeResponse = DecodeFinalizePasswordRecoveryFlowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildFinalizePasswordRecoveryFlowRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.FinalizePasswordRecoveryFlowDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "FinalizePasswordRecoveryFlow", err)
		}
		return decodeResponse(resp)
	}
}

// CreateOrder returns an endpoint that makes HTTP requests to the coachee
// service CreateOrder server.
func (c *Client) CreateOrder() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateOrderRequest(c.encoder)
		decodeResponse = DecodeCreateOrderResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateOrderRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateOrderDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "CreateOrder", err)
		}
		return decodeResponse(resp)
	}
}

// RegisterStripeExpress returns an endpoint that makes HTTP requests to the
// coachee service RegisterStripeExpress server.
func (c *Client) RegisterStripeExpress() goa.Endpoint {
	var (
		encodeRequest  = EncodeRegisterStripeExpressRequest(c.encoder)
		decodeResponse = DecodeRegisterStripeExpressResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRegisterStripeExpressRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RegisterStripeExpressDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "RegisterStripeExpress", err)
		}
		return decodeResponse(resp)
	}
}

// AdminLogin returns an endpoint that makes HTTP requests to the coachee
// service AdminLogin server.
func (c *Client) AdminLogin() goa.Endpoint {
	var (
		encodeRequest  = EncodeAdminLoginRequest(c.encoder)
		decodeResponse = DecodeAdminLoginResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAdminLoginRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AdminLoginDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "AdminLogin", err)
		}
		return decodeResponse(resp)
	}
}

// RegisterNewsletterEmail returns an endpoint that makes HTTP requests to the
// coachee service RegisterNewsletterEmail server.
func (c *Client) RegisterNewsletterEmail() goa.Endpoint {
	var (
		encodeRequest  = EncodeRegisterNewsletterEmailRequest(c.encoder)
		decodeResponse = DecodeRegisterNewsletterEmailResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRegisterNewsletterEmailRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RegisterNewsletterEmailDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("coachee", "RegisterNewsletterEmail", err)
		}
		return decodeResponse(resp)
	}
}
