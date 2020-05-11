// Code generated by goa v3.0.9, DO NOT EDIT.
//
// coachee service
//
// Command:
// $ goa gen coachee-backend/design

package coachee

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// The coachee service performs operations on coachees
type Service interface {
	// Stripe webhook endpoint
	StripeWebhooks(context.Context, string) (err error)
	// GetCoaches returns an array of coaches according to a tag and pagination
	GetCoaches(context.Context, *GetCoachesPayload) (res []*Coach, err error)
	// GetCoach returns one coach according to the id
	GetCoach(context.Context, *GetCoachPayload) (res *Coach, err error)
	// AdminGetCoach returns all coach info according to the id
	AdminGetCoach(context.Context, *AdminGetCoachPayload) (res *FullCoach, err error)
	// LenCoaches returns the amount of coaches with a given tag
	LenCoaches(context.Context, *LenCoachesPayload) (res uint, err error)
	// CreateCoaches creates a base coach
	CreateCoach(context.Context, *CreateCoachPayload) (res uint, err error)
	// Logs in a coach to stripe express
	LoginCoach(context.Context, *LoginCoachPayload) (res *LoginCoachResult, err error)
	// starts the process of recovering a password
	StartCoachPasswordRecoveryFlow(context.Context, *StartCoachPasswordRecoveryFlowPayload) (err error)
	// verifies if a recovery token is still valid
	CheckCoachPasswordRecoveryToken(context.Context, *CheckCoachPasswordRecoveryTokenPayload) (err error)
	// finalizes the password recovery flow by resetting a new password
	FinalizeCoachPasswordRecoveryFlow(context.Context, *FinalizeCoachPasswordRecoveryFlowPayload) (err error)
	// UpdateCoaches updates a coach
	UpdateCoach(context.Context, *UpdateCoachPayload) (err error)
	// creates a certification for a coach
	CreateCertification(context.Context, *CreateCertificationPayload) (err error)
	// deletes a certification for a coach
	DeleteCertification(context.Context, *DeleteCertificationPayload) (err error)
	// creates a program for a coach
	CreateProgram(context.Context, *CreateProgramPayload) (err error)
	// deletes a program for a coach
	DeleteProgram(context.Context, *DeleteProgramPayload) (err error)
	// creates an availability for a coach
	CreateAvailability(context.Context, *CreateAvailabilityPayload) (err error)
	// deletes an availability for a coach
	DeleteAvailability(context.Context, *DeleteAvailabilityPayload) (err error)
	// creates a new customer
	CreateCustomer(context.Context, *CreateCustomerPayload) (res *CreateCustomerResult, err error)
	// logs in a customer and returns a jwt
	CustomerLogin(context.Context, *CustomerLoginPayload) (res *CustomerLoginResult, err error)
	// starts the process of recovering a password
	StartPasswordRecoveryFlow(context.Context, *StartPasswordRecoveryFlowPayload) (err error)
	// verifies if a recovery token is still valid
	CheckPasswordRecoveryToken(context.Context, *CheckPasswordRecoveryTokenPayload) (err error)
	// finalizes the password recovery flow by resetting a new password
	FinalizePasswordRecoveryFlow(context.Context, *FinalizePasswordRecoveryFlowPayload) (err error)
	// creates a new order
	CreateOrder(context.Context, *CreateOrderPayload) (res *CreateOrderResult, err error)
	// registers a stripe express account in stripe and associates it to a coach
	RegisterStripeExpress(context.Context, *RegisterStripeExpressPayload) (err error)
	// logs in a customer and returns a jwt
	AdminLogin(context.Context, *AdminLoginPayload) (res *AdminLoginResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "coachee"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [25]string{"StripeWebhooks", "GetCoaches", "GetCoach", "AdminGetCoach", "LenCoaches", "CreateCoach", "LoginCoach", "StartCoachPasswordRecoveryFlow", "CheckCoachPasswordRecoveryToken", "FinalizeCoachPasswordRecoveryFlow", "UpdateCoach", "CreateCertification", "DeleteCertification", "CreateProgram", "DeleteProgram", "CreateAvailability", "DeleteAvailability", "CreateCustomer", "CustomerLogin", "StartPasswordRecoveryFlow", "CheckPasswordRecoveryToken", "FinalizePasswordRecoveryFlow", "CreateOrder", "RegisterStripeExpress", "AdminLogin"}

// GetCoachesPayload is the payload type of the coachee service GetCoaches
// method.
type GetCoachesPayload struct {
	Tag     *string
	Limit   *uint
	Page    *uint
	ShowAll *bool
}

// GetCoachPayload is the payload type of the coachee service GetCoach method.
type GetCoachPayload struct {
	ID uint
}

// Coach is the result type of the coachee service GetCoach method.
type Coach struct {
	ID             uint
	FirstName      string
	LastName       string
	Tags           string
	Description    string
	City           string
	Country        string
	PictureURL     string
	Certifications []*Certification
	Programs       []*Program
	Availability   []*Availability
}

// AdminGetCoachPayload is the payload type of the coachee service
// AdminGetCoach method.
type AdminGetCoachPayload struct {
	// JWT token used to perform authorization
	Token string
	ID    uint
}

// FullCoach is the result type of the coachee service AdminGetCoach method.
type FullCoach struct {
	ID             uint
	FirstName      string
	LastName       string
	Email          string
	Phone          string
	StripeID       string
	Tags           string
	Description    string
	City           string
	Country        string
	PictureURL     string
	Status         string
	Vat            string
	IntroCall      int
	Availability   []*Availability
	Certifications []*Certification
	Programs       []*Program
}

// LenCoachesPayload is the payload type of the coachee service LenCoaches
// method.
type LenCoachesPayload struct {
	Tag string
}

// CreateCoachPayload is the payload type of the coachee service CreateCoach
// method.
type CreateCoachPayload struct {
	FirstName          string
	LastName           string
	Email              string
	Password           string
	Phone              string
	Tags               string
	Description        string
	City               *string
	Country            *string
	IntroCall          uint
	TextCertifications string
	TextPrograms       string
	TextAvailability   *string
	Vat                *string
}

// LoginCoachPayload is the payload type of the coachee service LoginCoach
// method.
type LoginCoachPayload struct {
	Email    string
	Password string
}

// LoginCoachResult is the result type of the coachee service LoginCoach method.
type LoginCoachResult struct {
	URL string
}

// StartCoachPasswordRecoveryFlowPayload is the payload type of the coachee
// service StartCoachPasswordRecoveryFlow method.
type StartCoachPasswordRecoveryFlowPayload struct {
	Email string
}

// CheckCoachPasswordRecoveryTokenPayload is the payload type of the coachee
// service CheckCoachPasswordRecoveryToken method.
type CheckCoachPasswordRecoveryTokenPayload struct {
	Token string
}

// FinalizeCoachPasswordRecoveryFlowPayload is the payload type of the coachee
// service FinalizeCoachPasswordRecoveryFlow method.
type FinalizeCoachPasswordRecoveryFlowPayload struct {
	Token    string
	Password string
}

// UpdateCoachPayload is the payload type of the coachee service UpdateCoach
// method.
type UpdateCoachPayload struct {
	// JWT token used to perform authorization
	Token       string
	ID          uint
	FirstName   *string
	LastName    *string
	Email       *string
	Phone       *string
	Tags        *string
	Description *string
	City        *string
	Country     *string
	IntroCall   *uint
	StripeID    *string
	PictureURL  *string
	Vat         *string
	Status      *string
}

// CreateCertificationPayload is the payload type of the coachee service
// CreateCertification method.
type CreateCertificationPayload struct {
	// JWT token used to perform authorization
	Token         string
	ID            uint
	Certification *Certification
}

// DeleteCertificationPayload is the payload type of the coachee service
// DeleteCertification method.
type DeleteCertificationPayload struct {
	// JWT token used to perform authorization
	Token  string
	ID     uint
	CertID string
}

// CreateProgramPayload is the payload type of the coachee service
// CreateProgram method.
type CreateProgramPayload struct {
	// JWT token used to perform authorization
	Token   string
	ID      uint
	Program *Program
}

// DeleteProgramPayload is the payload type of the coachee service
// DeleteProgram method.
type DeleteProgramPayload struct {
	// JWT token used to perform authorization
	Token     string
	ID        uint
	ProgramID string
}

// represents a coach availability
type CreateAvailabilityPayload struct {
	// JWT token used to perform authorization
	Token   string
	ID      uint
	WeekDay uint
	Start   uint
	End     uint
}

// DeleteAvailabilityPayload is the payload type of the coachee service
// DeleteAvailability method.
type DeleteAvailabilityPayload struct {
	// JWT token used to perform authorization
	Token string
	ID    uint
	AvID  string
}

// CreateCustomerPayload is the payload type of the coachee service
// CreateCustomer method.
type CreateCustomerPayload struct {
	Email     string
	FirstName string
	LastName  string
	BirthDate int64
	Password  string
}

// CreateCustomerResult is the result type of the coachee service
// CreateCustomer method.
type CreateCustomerResult struct {
	Token  string
	Expiry int64
	User   *BaseClient
}

// CustomerLoginPayload is the payload type of the coachee service
// CustomerLogin method.
type CustomerLoginPayload struct {
	Email    string
	Password string
}

// CustomerLoginResult is the result type of the coachee service CustomerLogin
// method.
type CustomerLoginResult struct {
	Token  string
	Expiry int64
	User   *BaseClient
}

// StartPasswordRecoveryFlowPayload is the payload type of the coachee service
// StartPasswordRecoveryFlow method.
type StartPasswordRecoveryFlowPayload struct {
	Email string
}

// CheckPasswordRecoveryTokenPayload is the payload type of the coachee service
// CheckPasswordRecoveryToken method.
type CheckPasswordRecoveryTokenPayload struct {
	Token string
}

// FinalizePasswordRecoveryFlowPayload is the payload type of the coachee
// service FinalizePasswordRecoveryFlow method.
type FinalizePasswordRecoveryFlowPayload struct {
	Token    string
	Password string
}

// CreateOrderPayload is the payload type of the coachee service CreateOrder
// method.
type CreateOrderPayload struct {
	// JWT token used to perform authorization
	Token     string
	CoachID   uint
	ProgramID string
	IntroCall int64
}

// CreateOrderResult is the result type of the coachee service CreateOrder
// method.
type CreateOrderResult struct {
	ClientSecret  string
	PublishingKey string
}

// RegisterStripeExpressPayload is the payload type of the coachee service
// RegisterStripeExpress method.
type RegisterStripeExpressPayload struct {
	ID                uint
	AuthorizationCode string
}

// AdminLoginPayload is the payload type of the coachee service AdminLogin
// method.
type AdminLoginPayload struct {
	Email    string
	Password string
}

// AdminLoginResult is the result type of the coachee service AdminLogin method.
type AdminLoginResult struct {
	Token  string
	Expiry int64
}

// represents a coach certification
type Certification struct {
	ID          *string
	Title       string
	Description string
	Institution string
	Month       uint
	Year        uint
}

// represents a coach's programs
type Program struct {
	ID          *string
	Name        string
	Sessions    uint
	Duration    uint
	Description string
	TotalPrice  uint
	TaxPercent  uint
}

// represents a coach availability
type Availability struct {
	ID        string
	WeekDay   uint
	Start     float64
	End       float64
	DateLabel string
}

// represents a client
type BaseClient struct {
	ID        uint
	FirstName string
	LastName  string
}

// MakeTransient builds a goa.ServiceError from an error.
func MakeTransient(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:      "transient",
		ID:        goa.NewErrorID(),
		Message:   err.Error(),
		Temporary: true,
	}
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "notFound",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeValidation builds a goa.ServiceError from an error.
func MakeValidation(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "validation",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "unauthorized",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeInternal builds a goa.ServiceError from an error.
func MakeInternal(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "internal",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
