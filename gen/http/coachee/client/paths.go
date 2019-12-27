// Code generated by goa v3.0.7, DO NOT EDIT.
//
// HTTP request path constructors for the coachee service.
//
// Command:
// $ goa gen coachee-backend/design

package client

import (
	"fmt"
)

// GetCoachesCoacheePath returns the URL path to the coachee service GetCoaches HTTP endpoint.
func GetCoachesCoacheePath() string {
	return "/coaches"
}

// GetCoachCoacheePath returns the URL path to the coachee service GetCoach HTTP endpoint.
func GetCoachCoacheePath(id uint) string {
	return fmt.Sprintf("/coaches/%v", id)
}

// LenCoachesCoacheePath returns the URL path to the coachee service LenCoaches HTTP endpoint.
func LenCoachesCoacheePath(tag string) string {
	return fmt.Sprintf("/coaches/%v/length", tag)
}

// CreateCoachCoacheePath returns the URL path to the coachee service CreateCoach HTTP endpoint.
func CreateCoachCoacheePath() string {
	return "/coaches"
}

// UpdateCoachCoacheePath returns the URL path to the coachee service UpdateCoach HTTP endpoint.
func UpdateCoachCoacheePath(id uint) string {
	return fmt.Sprintf("/coaches/%v", id)
}

// CreateCertificationCoacheePath returns the URL path to the coachee service CreateCertification HTTP endpoint.
func CreateCertificationCoacheePath(id uint) string {
	return fmt.Sprintf("/coaches/%v/certifications", id)
}

// DeleteCertificationCoacheePath returns the URL path to the coachee service DeleteCertification HTTP endpoint.
func DeleteCertificationCoacheePath(id uint, certID string) string {
	return fmt.Sprintf("/coaches/%v/certifications/%v", id, certID)
}

// CreateProgramCoacheePath returns the URL path to the coachee service CreateProgram HTTP endpoint.
func CreateProgramCoacheePath(id uint) string {
	return fmt.Sprintf("/coaches/%v/programs", id)
}

// DeleteProgramCoacheePath returns the URL path to the coachee service DeleteProgram HTTP endpoint.
func DeleteProgramCoacheePath(id uint, programID string) string {
	return fmt.Sprintf("/coaches/%v/programs/%v", id, programID)
}

// CreateAvailabilityCoacheePath returns the URL path to the coachee service CreateAvailability HTTP endpoint.
func CreateAvailabilityCoacheePath(id uint) string {
	return fmt.Sprintf("/coaches/%v/availability", id)
}

// DeleteAvailabilityCoacheePath returns the URL path to the coachee service DeleteAvailability HTTP endpoint.
func DeleteAvailabilityCoacheePath(id uint, avID string) string {
	return fmt.Sprintf("/coaches/%v/availability/%v", id, avID)
}

// CreateCustomerCoacheePath returns the URL path to the coachee service CreateCustomer HTTP endpoint.
func CreateCustomerCoacheePath() string {
	return "/clients"
}

// CustomerLoginCoacheePath returns the URL path to the coachee service CustomerLogin HTTP endpoint.
func CustomerLoginCoacheePath() string {
	return "/clients/login"
}

// CreateOrderCoacheePath returns the URL path to the coachee service CreateOrder HTTP endpoint.
func CreateOrderCoacheePath() string {
	return "/orders"
}
