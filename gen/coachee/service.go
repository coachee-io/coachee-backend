// Code generated by goa v3.0.7, DO NOT EDIT.
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
	// GetCoaches returns an array of coaches according to a tag and pagination
	GetCoaches(context.Context, *GetCoachesPayload) (res []*Coach, err error)
	// GetCoach returns one coach according to the id
	GetCoach(context.Context, *GetCoachPayload) (res *Coach, err error)
	// LenCoaches returns the amount of coaches with a given tag
	LenCoaches(context.Context, *LenCoachesPayload) (res uint, err error)
	// CreateCoaches creates a base coach
	CreateCoach(context.Context, *CreateCoachPayload) (res uint, err error)
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
	// creates a new client
	CreateClient(context.Context, *CreateClientPayload) (res string, err error)
	// ClientLogin implements ClientLogin.
	ClientLogin(context.Context, *ClientLoginPayload) (res string, err error)
	// CreateOrder implements CreateOrder.
	CreateOrder(context.Context, *CreateOrderPayload) (err error)
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
var MethodNames = [14]string{"GetCoaches", "GetCoach", "LenCoaches", "CreateCoach", "UpdateCoach", "CreateCertification", "DeleteCertification", "CreateProgram", "DeleteProgram", "CreateAvailability", "DeleteAvailability", "CreateClient", "ClientLogin", "CreateOrder"}

// GetCoachesPayload is the payload type of the coachee service GetCoaches
// method.
type GetCoachesPayload struct {
	Tag   *string
	Limit *uint
	Page  *uint
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

// UpdateCoachPayload is the payload type of the coachee service UpdateCoach
// method.
type UpdateCoachPayload struct {
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
}

// CreateCertificationPayload is the payload type of the coachee service
// CreateCertification method.
type CreateCertificationPayload struct {
	ID            uint
	Certification *Certification
}

// DeleteCertificationPayload is the payload type of the coachee service
// DeleteCertification method.
type DeleteCertificationPayload struct {
	ID     uint
	CertID string
}

// CreateProgramPayload is the payload type of the coachee service
// CreateProgram method.
type CreateProgramPayload struct {
	ID      uint
	Program *Program
}

// DeleteProgramPayload is the payload type of the coachee service
// DeleteProgram method.
type DeleteProgramPayload struct {
	ID        uint
	ProgramID string
}

// CreateAvailabilityPayload is the payload type of the coachee service
// CreateAvailability method.
type CreateAvailabilityPayload struct {
	ID           uint
	Availability *Availability
}

// DeleteAvailabilityPayload is the payload type of the coachee service
// DeleteAvailability method.
type DeleteAvailabilityPayload struct {
	ID   uint
	AvID string
}

// CreateClientPayload is the payload type of the coachee service CreateClient
// method.
type CreateClientPayload struct {
	Email     string
	FirstName string
	LastName  string
	BirthDate int64
	Password  string
}

// ClientLoginPayload is the payload type of the coachee service ClientLogin
// method.
type ClientLoginPayload struct {
	Email    string
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
	ID      *string
	WeekDay uint
	Start   uint
	End     uint
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
