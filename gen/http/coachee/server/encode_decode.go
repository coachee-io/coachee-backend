// Code generated by goa v3.0.7, DO NOT EDIT.
//
// coachee HTTP server encoders and decoders
//
// Command:
// $ goa gen coachee-backend/design

package server

import (
	coachee "coachee-backend/gen/coachee"
	"context"
	"io"
	"net/http"
	"strconv"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetCoachesResponse returns an encoder for responses returned by the
// coachee GetCoaches endpoint.
func EncodeGetCoachesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]*coachee.Coach)
		enc := encoder(ctx, w)
		body := NewGetCoachesResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetCoachesRequest returns a decoder for requests sent to the coachee
// GetCoaches endpoint.
func DecodeGetCoachesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			tag   string
			limit *uint
			page  *uint
			err   error

			params = mux.Vars(r)
		)
		tag = params["tag"]
		{
			limitRaw := r.URL.Query().Get("limit")
			if limitRaw != "" {
				v, err2 := strconv.ParseUint(limitRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("limit", limitRaw, "unsigned integer"))
				}
				pv := uint(v)
				limit = &pv
			}
		}
		{
			pageRaw := r.URL.Query().Get("page")
			if pageRaw != "" {
				v, err2 := strconv.ParseUint(pageRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("page", pageRaw, "unsigned integer"))
				}
				pv := uint(v)
				page = &pv
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetCoachesPayload(tag, limit, page)

		return payload, nil
	}
}

// EncodeGetCoachesError returns an encoder for errors returned by the
// GetCoaches coachee endpoint.
func EncodeGetCoachesError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "transient":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewGetCoachesTransientResponseBody(res)
			w.Header().Set("goa-error", "transient")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "notFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewGetCoachesNotFoundResponseBody(res)
			w.Header().Set("goa-error", "notFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "validation":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewGetCoachesValidationResponseBody(res)
			w.Header().Set("goa-error", "validation")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewGetCoachesUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeLenCoachesResponse returns an encoder for responses returned by the
// coachee LenCoaches endpoint.
func EncodeLenCoachesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(uint)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeLenCoachesRequest returns a decoder for requests sent to the coachee
// LenCoaches endpoint.
func DecodeLenCoachesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			tag string

			params = mux.Vars(r)
		)
		tag = params["tag"]
		payload := NewLenCoachesPayload(tag)

		return payload, nil
	}
}

// EncodeLenCoachesError returns an encoder for errors returned by the
// LenCoaches coachee endpoint.
func EncodeLenCoachesError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "transient":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewLenCoachesTransientResponseBody(res)
			w.Header().Set("goa-error", "transient")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "notFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewLenCoachesNotFoundResponseBody(res)
			w.Header().Set("goa-error", "notFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "validation":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewLenCoachesValidationResponseBody(res)
			w.Header().Set("goa-error", "validation")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewLenCoachesUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateCoachResponse returns an encoder for responses returned by the
// coachee CreateCoach endpoint.
func EncodeCreateCoachResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(uint)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateCoachRequest returns a decoder for requests sent to the coachee
// CreateCoach endpoint.
func DecodeCreateCoachRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateCoachRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateCoachRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateCoachPayload(&body)

		return payload, nil
	}
}

// EncodeCreateCoachError returns an encoder for errors returned by the
// CreateCoach coachee endpoint.
func EncodeCreateCoachError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "transient":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateCoachTransientResponseBody(res)
			w.Header().Set("goa-error", "transient")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "notFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateCoachNotFoundResponseBody(res)
			w.Header().Set("goa-error", "notFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "validation":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateCoachValidationResponseBody(res)
			w.Header().Set("goa-error", "validation")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateCoachUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateCoachResponse returns an encoder for responses returned by the
// coachee UpdateCoach endpoint.
func EncodeUpdateCoachResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusAccepted)
		return nil
	}
}

// DecodeUpdateCoachRequest returns a decoder for requests sent to the coachee
// UpdateCoach endpoint.
func DecodeUpdateCoachRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateCoachRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			id uint

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdateCoachPayload(&body, id)

		return payload, nil
	}
}

// EncodeUpdateCoachError returns an encoder for errors returned by the
// UpdateCoach coachee endpoint.
func EncodeUpdateCoachError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "transient":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewUpdateCoachTransientResponseBody(res)
			w.Header().Set("goa-error", "transient")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "notFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewUpdateCoachNotFoundResponseBody(res)
			w.Header().Set("goa-error", "notFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "validation":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewUpdateCoachValidationResponseBody(res)
			w.Header().Set("goa-error", "validation")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewUpdateCoachUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateCertificationResponse returns an encoder for responses returned
// by the coachee CreateCertification endpoint.
func EncodeCreateCertificationResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusAccepted)
		return nil
	}
}

// DecodeCreateCertificationRequest returns a decoder for requests sent to the
// coachee CreateCertification endpoint.
func DecodeCreateCertificationRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateCertificationRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateCertificationRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id uint

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewCreateCertificationPayload(&body, id)

		return payload, nil
	}
}

// EncodeCreateCertificationError returns an encoder for errors returned by the
// CreateCertification coachee endpoint.
func EncodeCreateCertificationError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "transient":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateCertificationTransientResponseBody(res)
			w.Header().Set("goa-error", "transient")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "notFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateCertificationNotFoundResponseBody(res)
			w.Header().Set("goa-error", "notFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "validation":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateCertificationValidationResponseBody(res)
			w.Header().Set("goa-error", "validation")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateCertificationUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteCertificationResponse returns an encoder for responses returned
// by the coachee DeleteCertification endpoint.
func EncodeDeleteCertificationResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDeleteCertificationRequest returns a decoder for requests sent to the
// coachee DeleteCertification endpoint.
func DecodeDeleteCertificationRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id     uint
			certID string
			err    error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		certID = params["certID"]
		if err != nil {
			return nil, err
		}
		payload := NewDeleteCertificationPayload(id, certID)

		return payload, nil
	}
}

// EncodeDeleteCertificationError returns an encoder for errors returned by the
// DeleteCertification coachee endpoint.
func EncodeDeleteCertificationError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "transient":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteCertificationTransientResponseBody(res)
			w.Header().Set("goa-error", "transient")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "notFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteCertificationNotFoundResponseBody(res)
			w.Header().Set("goa-error", "notFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "validation":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteCertificationValidationResponseBody(res)
			w.Header().Set("goa-error", "validation")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteCertificationUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateProgramResponse returns an encoder for responses returned by the
// coachee CreateProgram endpoint.
func EncodeCreateProgramResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusAccepted)
		return nil
	}
}

// DecodeCreateProgramRequest returns a decoder for requests sent to the
// coachee CreateProgram endpoint.
func DecodeCreateProgramRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateProgramRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateProgramRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id uint

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewCreateProgramPayload(&body, id)

		return payload, nil
	}
}

// EncodeCreateProgramError returns an encoder for errors returned by the
// CreateProgram coachee endpoint.
func EncodeCreateProgramError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "transient":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateProgramTransientResponseBody(res)
			w.Header().Set("goa-error", "transient")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "notFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateProgramNotFoundResponseBody(res)
			w.Header().Set("goa-error", "notFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "validation":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateProgramValidationResponseBody(res)
			w.Header().Set("goa-error", "validation")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateProgramUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteProgramResponse returns an encoder for responses returned by the
// coachee DeleteProgram endpoint.
func EncodeDeleteProgramResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDeleteProgramRequest returns a decoder for requests sent to the
// coachee DeleteProgram endpoint.
func DecodeDeleteProgramRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id        uint
			programID string
			err       error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		programID = params["programID"]
		if err != nil {
			return nil, err
		}
		payload := NewDeleteProgramPayload(id, programID)

		return payload, nil
	}
}

// EncodeDeleteProgramError returns an encoder for errors returned by the
// DeleteProgram coachee endpoint.
func EncodeDeleteProgramError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "transient":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteProgramTransientResponseBody(res)
			w.Header().Set("goa-error", "transient")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "notFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteProgramNotFoundResponseBody(res)
			w.Header().Set("goa-error", "notFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "validation":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteProgramValidationResponseBody(res)
			w.Header().Set("goa-error", "validation")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteProgramUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateAvailabilityResponse returns an encoder for responses returned
// by the coachee CreateAvailability endpoint.
func EncodeCreateAvailabilityResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusAccepted)
		return nil
	}
}

// DecodeCreateAvailabilityRequest returns a decoder for requests sent to the
// coachee CreateAvailability endpoint.
func DecodeCreateAvailabilityRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateAvailabilityRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateAvailabilityRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id uint

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewCreateAvailabilityPayload(&body, id)

		return payload, nil
	}
}

// EncodeCreateAvailabilityError returns an encoder for errors returned by the
// CreateAvailability coachee endpoint.
func EncodeCreateAvailabilityError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "transient":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateAvailabilityTransientResponseBody(res)
			w.Header().Set("goa-error", "transient")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "notFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateAvailabilityNotFoundResponseBody(res)
			w.Header().Set("goa-error", "notFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "validation":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateAvailabilityValidationResponseBody(res)
			w.Header().Set("goa-error", "validation")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateAvailabilityUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteAvailabilityResponse returns an encoder for responses returned
// by the coachee DeleteAvailability endpoint.
func EncodeDeleteAvailabilityResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDeleteAvailabilityRequest returns a decoder for requests sent to the
// coachee DeleteAvailability endpoint.
func DecodeDeleteAvailabilityRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id   uint
			avID string
			err  error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		avID = params["avID"]
		if err != nil {
			return nil, err
		}
		payload := NewDeleteAvailabilityPayload(id, avID)

		return payload, nil
	}
}

// EncodeDeleteAvailabilityError returns an encoder for errors returned by the
// DeleteAvailability coachee endpoint.
func EncodeDeleteAvailabilityError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "transient":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteAvailabilityTransientResponseBody(res)
			w.Header().Set("goa-error", "transient")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "notFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteAvailabilityNotFoundResponseBody(res)
			w.Header().Set("goa-error", "notFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "validation":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteAvailabilityValidationResponseBody(res)
			w.Header().Set("goa-error", "validation")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteAvailabilityUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalCoacheeCertificationToCertificationResponse builds a value of type
// *CertificationResponse from a value of type *coachee.Certification.
func marshalCoacheeCertificationToCertificationResponse(v *coachee.Certification) *CertificationResponse {
	if v == nil {
		return nil
	}
	res := &CertificationResponse{
		ID:          v.ID,
		Title:       v.Title,
		Description: v.Description,
		Institution: v.Institution,
		Month:       v.Month,
		Year:        v.Year,
	}

	return res
}

// marshalCoacheeProgramToProgramResponse builds a value of type
// *ProgramResponse from a value of type *coachee.Program.
func marshalCoacheeProgramToProgramResponse(v *coachee.Program) *ProgramResponse {
	if v == nil {
		return nil
	}
	res := &ProgramResponse{
		ID:          v.ID,
		Name:        v.Name,
		Sessions:    v.Sessions,
		Duration:    v.Duration,
		Description: v.Description,
		TotalPrice:  v.TotalPrice,
		TaxPercent:  v.TaxPercent,
	}

	return res
}

// marshalCoacheeAvailabilityToAvailabilityResponse builds a value of type
// *AvailabilityResponse from a value of type *coachee.Availability.
func marshalCoacheeAvailabilityToAvailabilityResponse(v *coachee.Availability) *AvailabilityResponse {
	if v == nil {
		return nil
	}
	res := &AvailabilityResponse{
		ID:      v.ID,
		WeekDay: v.WeekDay,
		Start:   v.Start,
		End:     v.End,
	}

	return res
}

// unmarshalCertificationRequestBodyToCoacheeCertification builds a value of
// type *coachee.Certification from a value of type *CertificationRequestBody.
func unmarshalCertificationRequestBodyToCoacheeCertification(v *CertificationRequestBody) *coachee.Certification {
	res := &coachee.Certification{
		ID:          v.ID,
		Title:       *v.Title,
		Description: *v.Description,
		Institution: *v.Institution,
		Month:       *v.Month,
		Year:        *v.Year,
	}

	return res
}

// unmarshalProgramRequestBodyToCoacheeProgram builds a value of type
// *coachee.Program from a value of type *ProgramRequestBody.
func unmarshalProgramRequestBodyToCoacheeProgram(v *ProgramRequestBody) *coachee.Program {
	res := &coachee.Program{
		ID:          v.ID,
		Name:        *v.Name,
		Sessions:    *v.Sessions,
		Duration:    *v.Duration,
		Description: *v.Description,
		TotalPrice:  *v.TotalPrice,
		TaxPercent:  *v.TaxPercent,
	}

	return res
}

// unmarshalAvailabilityRequestBodyToCoacheeAvailability builds a value of type
// *coachee.Availability from a value of type *AvailabilityRequestBody.
func unmarshalAvailabilityRequestBodyToCoacheeAvailability(v *AvailabilityRequestBody) *coachee.Availability {
	res := &coachee.Availability{
		ID:      v.ID,
		WeekDay: *v.WeekDay,
		Start:   *v.Start,
		End:     *v.End,
	}

	return res
}
