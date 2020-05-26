// Code generated by goa v3.0.9, DO NOT EDIT.
//
// coachee HTTP client CLI support package
//
// Command:
// $ goa gen coachee-backend/design

package client

import (
	coachee "coachee-backend/gen/coachee"
	"encoding/json"
	"fmt"
	"strconv"

	goa "goa.design/goa/v3/pkg"
)

// BuildGetCoachesPayload builds the payload for the coachee GetCoaches
// endpoint from CLI flags.
func BuildGetCoachesPayload(coacheeGetCoachesTag string, coacheeGetCoachesLimit string, coacheeGetCoachesPage string, coacheeGetCoachesShowAll string) (*coachee.GetCoachesPayload, error) {
	var err error
	var tag *string
	{
		if coacheeGetCoachesTag != "" {
			tag = &coacheeGetCoachesTag
		}
	}
	var limit *uint
	{
		if coacheeGetCoachesLimit != "" {
			var v uint64
			v, err = strconv.ParseUint(coacheeGetCoachesLimit, 10, 64)
			val := uint(v)
			limit = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be UINT")
			}
		}
	}
	var page *uint
	{
		if coacheeGetCoachesPage != "" {
			var v uint64
			v, err = strconv.ParseUint(coacheeGetCoachesPage, 10, 64)
			val := uint(v)
			page = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for page, must be UINT")
			}
		}
	}
	var showAll *bool
	{
		if coacheeGetCoachesShowAll != "" {
			var val bool
			val, err = strconv.ParseBool(coacheeGetCoachesShowAll)
			showAll = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for showAll, must be BOOL")
			}
		}
	}
	payload := &coachee.GetCoachesPayload{
		Tag:     tag,
		Limit:   limit,
		Page:    page,
		ShowAll: showAll,
	}
	return payload, nil
}

// BuildGetCoachPayload builds the payload for the coachee GetCoach endpoint
// from CLI flags.
func BuildGetCoachPayload(coacheeGetCoachID string) (*coachee.GetCoachPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(coacheeGetCoachID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	payload := &coachee.GetCoachPayload{
		ID: id,
	}
	return payload, nil
}

// BuildAdminGetCoachPayload builds the payload for the coachee AdminGetCoach
// endpoint from CLI flags.
func BuildAdminGetCoachPayload(coacheeAdminGetCoachID string, coacheeAdminGetCoachToken string) (*coachee.AdminGetCoachPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(coacheeAdminGetCoachID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	var token string
	{
		token = coacheeAdminGetCoachToken
	}
	payload := &coachee.AdminGetCoachPayload{
		ID:    id,
		Token: token,
	}
	return payload, nil
}

// BuildLenCoachesPayload builds the payload for the coachee LenCoaches
// endpoint from CLI flags.
func BuildLenCoachesPayload(coacheeLenCoachesTag string) (*coachee.LenCoachesPayload, error) {
	var tag string
	{
		tag = coacheeLenCoachesTag
	}
	payload := &coachee.LenCoachesPayload{
		Tag: tag,
	}
	return payload, nil
}

// BuildCreateCoachPayload builds the payload for the coachee CreateCoach
// endpoint from CLI flags.
func BuildCreateCoachPayload(coacheeCreateCoachBody string) (*coachee.CreateCoachPayload, error) {
	var err error
	var body CreateCoachRequestBody
	{
		err = json.Unmarshal([]byte(coacheeCreateCoachBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"city\": \"In enim quo commodi quia nihil.\",\n      \"country\": \"Eos ut est nam earum illum eos.\",\n      \"description\": \"Autem fugiat et.\",\n      \"email\": \"Laboriosam tempora modi.\",\n      \"firstName\": \"Culpa quo repellat.\",\n      \"introCall\": 6289881553078306697,\n      \"lastName\": \"Assumenda ratione sit dolor qui et sint.\",\n      \"password\": \"Quos eum rerum architecto et facere.\",\n      \"phone\": \"Aut libero rerum mollitia accusantium tempore fugit.\",\n      \"tags\": \"Quod veritatis ad deleniti ut quisquam.\",\n      \"textAvailability\": \"Vitae sunt doloremque veritatis vel provident voluptas.\",\n      \"textCertifications\": \"Sint laborum excepturi eum repudiandae.\",\n      \"textPrograms\": \"Neque provident nemo.\",\n      \"vat\": \"Nisi maxime.\"\n   }'")
		}
	}
	v := &coachee.CreateCoachPayload{
		FirstName:          body.FirstName,
		LastName:           body.LastName,
		Email:              body.Email,
		Password:           body.Password,
		Phone:              body.Phone,
		Tags:               body.Tags,
		Description:        body.Description,
		City:               body.City,
		Country:            body.Country,
		IntroCall:          body.IntroCall,
		TextCertifications: body.TextCertifications,
		TextPrograms:       body.TextPrograms,
		TextAvailability:   body.TextAvailability,
		Vat:                body.Vat,
	}
	return v, nil
}

// BuildLoginCoachPayload builds the payload for the coachee LoginCoach
// endpoint from CLI flags.
func BuildLoginCoachPayload(coacheeLoginCoachBody string) (*coachee.LoginCoachPayload, error) {
	var err error
	var body LoginCoachRequestBody
	{
		err = json.Unmarshal([]byte(coacheeLoginCoachBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"email\": \"Iusto illo sit voluptatem dolor qui omnis.\",\n      \"password\": \"In omnis et.\"\n   }'")
		}
	}
	v := &coachee.LoginCoachPayload{
		Email:    body.Email,
		Password: body.Password,
	}
	return v, nil
}

// BuildStartCoachPasswordRecoveryFlowPayload builds the payload for the
// coachee StartCoachPasswordRecoveryFlow endpoint from CLI flags.
func BuildStartCoachPasswordRecoveryFlowPayload(coacheeStartCoachPasswordRecoveryFlowBody string) (*coachee.StartCoachPasswordRecoveryFlowPayload, error) {
	var err error
	var body StartCoachPasswordRecoveryFlowRequestBody
	{
		err = json.Unmarshal([]byte(coacheeStartCoachPasswordRecoveryFlowBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"email\": \"Sunt blanditiis illum in.\"\n   }'")
		}
	}
	v := &coachee.StartCoachPasswordRecoveryFlowPayload{
		Email: body.Email,
	}
	return v, nil
}

// BuildCheckCoachPasswordRecoveryTokenPayload builds the payload for the
// coachee CheckCoachPasswordRecoveryToken endpoint from CLI flags.
func BuildCheckCoachPasswordRecoveryTokenPayload(coacheeCheckCoachPasswordRecoveryTokenToken string) (*coachee.CheckCoachPasswordRecoveryTokenPayload, error) {
	var token string
	{
		token = coacheeCheckCoachPasswordRecoveryTokenToken
	}
	payload := &coachee.CheckCoachPasswordRecoveryTokenPayload{
		Token: token,
	}
	return payload, nil
}

// BuildFinalizeCoachPasswordRecoveryFlowPayload builds the payload for the
// coachee FinalizeCoachPasswordRecoveryFlow endpoint from CLI flags.
func BuildFinalizeCoachPasswordRecoveryFlowPayload(coacheeFinalizeCoachPasswordRecoveryFlowBody string, coacheeFinalizeCoachPasswordRecoveryFlowToken string) (*coachee.FinalizeCoachPasswordRecoveryFlowPayload, error) {
	var err error
	var body FinalizeCoachPasswordRecoveryFlowRequestBody
	{
		err = json.Unmarshal([]byte(coacheeFinalizeCoachPasswordRecoveryFlowBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"password\": \"Voluptatem vel rerum eum praesentium eum.\"\n   }'")
		}
	}
	var token string
	{
		token = coacheeFinalizeCoachPasswordRecoveryFlowToken
	}
	v := &coachee.FinalizeCoachPasswordRecoveryFlowPayload{
		Password: body.Password,
	}
	v.Token = token
	return v, nil
}

// BuildUpdateCoachPayload builds the payload for the coachee UpdateCoach
// endpoint from CLI flags.
func BuildUpdateCoachPayload(coacheeUpdateCoachBody string, coacheeUpdateCoachID string, coacheeUpdateCoachToken string) (*coachee.UpdateCoachPayload, error) {
	var err error
	var body UpdateCoachRequestBody
	{
		err = json.Unmarshal([]byte(coacheeUpdateCoachBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"city\": \"Non explicabo veritatis ea natus non at.\",\n      \"country\": \"Mollitia ea.\",\n      \"description\": \"Ut et dolores et expedita.\",\n      \"email\": \"Aut provident qui et quibusdam quod molestias.\",\n      \"firstCallDuration\": 1439109767,\n      \"firstName\": \"Et consequatur fuga nemo quaerat quia.\",\n      \"introCall\": 7449579427835838020,\n      \"lastName\": \"Aut aut ex.\",\n      \"phone\": \"Commodi deleniti ea ut.\",\n      \"pictureURL\": \"Qui tempora consequuntur excepturi delectus rerum.\",\n      \"status\": \"Repellat quisquam at quam doloribus nisi id.\",\n      \"stripeID\": \"Laudantium non eos.\",\n      \"tags\": \"Dicta dolorem omnis rerum dolorem.\",\n      \"vat\": \"Aut distinctio.\"\n   }'")
		}
	}
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(coacheeUpdateCoachID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	var token string
	{
		token = coacheeUpdateCoachToken
	}
	v := &coachee.UpdateCoachPayload{
		FirstName:         body.FirstName,
		LastName:          body.LastName,
		Email:             body.Email,
		Phone:             body.Phone,
		Tags:              body.Tags,
		Description:       body.Description,
		City:              body.City,
		Country:           body.Country,
		IntroCall:         body.IntroCall,
		StripeID:          body.StripeID,
		PictureURL:        body.PictureURL,
		Vat:               body.Vat,
		Status:            body.Status,
		FirstCallDuration: body.FirstCallDuration,
	}
	v.ID = id
	v.Token = token
	return v, nil
}

// BuildCreateCertificationPayload builds the payload for the coachee
// CreateCertification endpoint from CLI flags.
func BuildCreateCertificationPayload(coacheeCreateCertificationBody string, coacheeCreateCertificationID string, coacheeCreateCertificationToken string) (*coachee.CreateCertificationPayload, error) {
	var err error
	var body CreateCertificationRequestBody
	{
		err = json.Unmarshal([]byte(coacheeCreateCertificationBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"certification\": {\n         \"description\": \"Aut praesentium sint.\",\n         \"id\": \"Et id vel totam ipsam est magni.\",\n         \"institution\": \"Sint dolor eum non quae itaque.\",\n         \"month\": 7,\n         \"title\": \"In aut illum.\",\n         \"year\": 2024\n      }\n   }'")
		}
		if body.Certification == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("certification", "body"))
		}
		if body.Certification != nil {
			if err2 := ValidateCertificationRequestBody(body.Certification); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(coacheeCreateCertificationID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	var token string
	{
		token = coacheeCreateCertificationToken
	}
	v := &coachee.CreateCertificationPayload{}
	if body.Certification != nil {
		v.Certification = marshalCertificationRequestBodyToCoacheeCertification(body.Certification)
	}
	v.ID = id
	v.Token = token
	return v, nil
}

// BuildDeleteCertificationPayload builds the payload for the coachee
// DeleteCertification endpoint from CLI flags.
func BuildDeleteCertificationPayload(coacheeDeleteCertificationID string, coacheeDeleteCertificationCertID string, coacheeDeleteCertificationToken string) (*coachee.DeleteCertificationPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(coacheeDeleteCertificationID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	var certID string
	{
		certID = coacheeDeleteCertificationCertID
	}
	var token string
	{
		token = coacheeDeleteCertificationToken
	}
	payload := &coachee.DeleteCertificationPayload{
		ID:     id,
		CertID: certID,
		Token:  token,
	}
	return payload, nil
}

// BuildCreateProgramPayload builds the payload for the coachee CreateProgram
// endpoint from CLI flags.
func BuildCreateProgramPayload(coacheeCreateProgramBody string, coacheeCreateProgramID string, coacheeCreateProgramToken string) (*coachee.CreateProgramPayload, error) {
	var err error
	var body CreateProgramRequestBody
	{
		err = json.Unmarshal([]byte(coacheeCreateProgramBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"program\": {\n         \"description\": \"Enim similique et.\",\n         \"duration\": 8302032262535934794,\n         \"id\": \"Nostrum odio quia ea et aut qui.\",\n         \"name\": \"Voluptatibus numquam nihil non consequuntur.\",\n         \"sessions\": 17958077144555001509,\n         \"taxPercent\": 13312637532326067827,\n         \"totalPrice\": 4216243859987989031\n      }\n   }'")
		}
		if body.Program == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("program", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(coacheeCreateProgramID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	var token string
	{
		token = coacheeCreateProgramToken
	}
	v := &coachee.CreateProgramPayload{}
	if body.Program != nil {
		v.Program = marshalProgramRequestBodyToCoacheeProgram(body.Program)
	}
	v.ID = id
	v.Token = token
	return v, nil
}

// BuildDeleteProgramPayload builds the payload for the coachee DeleteProgram
// endpoint from CLI flags.
func BuildDeleteProgramPayload(coacheeDeleteProgramID string, coacheeDeleteProgramProgramID string, coacheeDeleteProgramToken string) (*coachee.DeleteProgramPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(coacheeDeleteProgramID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	var programID string
	{
		programID = coacheeDeleteProgramProgramID
	}
	var token string
	{
		token = coacheeDeleteProgramToken
	}
	payload := &coachee.DeleteProgramPayload{
		ID:        id,
		ProgramID: programID,
		Token:     token,
	}
	return payload, nil
}

// BuildCreateAvailabilityPayload builds the payload for the coachee
// CreateAvailability endpoint from CLI flags.
func BuildCreateAvailabilityPayload(coacheeCreateAvailabilityBody string, coacheeCreateAvailabilityID string, coacheeCreateAvailabilityToken string) (*coachee.CreateAvailabilityPayload, error) {
	var err error
	var body CreateAvailabilityRequestBody
	{
		err = json.Unmarshal([]byte(coacheeCreateAvailabilityBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"end\": 1069,\n      \"start\": 1379,\n      \"weekDay\": 3\n   }'")
		}
		if body.WeekDay < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.weekDay", body.WeekDay, 0, true))
		}
		if body.WeekDay > 6 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.weekDay", body.WeekDay, 6, false))
		}
		if body.Start < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.start", body.Start, 0, true))
		}
		if body.Start > 1440 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.start", body.Start, 1440, false))
		}
		if body.End < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.end", body.End, 0, true))
		}
		if body.End > 1440 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.end", body.End, 1440, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(coacheeCreateAvailabilityID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	var token string
	{
		token = coacheeCreateAvailabilityToken
	}
	v := &coachee.CreateAvailabilityPayload{
		WeekDay: body.WeekDay,
		Start:   body.Start,
		End:     body.End,
	}
	v.ID = id
	v.Token = token
	return v, nil
}

// BuildDeleteAvailabilityPayload builds the payload for the coachee
// DeleteAvailability endpoint from CLI flags.
func BuildDeleteAvailabilityPayload(coacheeDeleteAvailabilityID string, coacheeDeleteAvailabilityAvID string, coacheeDeleteAvailabilityToken string) (*coachee.DeleteAvailabilityPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(coacheeDeleteAvailabilityID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	var avID string
	{
		avID = coacheeDeleteAvailabilityAvID
	}
	var token string
	{
		token = coacheeDeleteAvailabilityToken
	}
	payload := &coachee.DeleteAvailabilityPayload{
		ID:    id,
		AvID:  avID,
		Token: token,
	}
	return payload, nil
}

// BuildCreateCustomerPayload builds the payload for the coachee CreateCustomer
// endpoint from CLI flags.
func BuildCreateCustomerPayload(coacheeCreateCustomerBody string) (*coachee.CreateCustomerPayload, error) {
	var err error
	var body CreateCustomerRequestBody
	{
		err = json.Unmarshal([]byte(coacheeCreateCustomerBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"birthDate\": 3515380160441746509,\n      \"email\": \"Assumenda consequatur explicabo perspiciatis asperiores debitis quas.\",\n      \"firstName\": \"Eum nisi natus autem.\",\n      \"lastName\": \"Officiis labore quo reiciendis.\",\n      \"password\": \"Possimus est inventore.\"\n   }'")
		}
	}
	v := &coachee.CreateCustomerPayload{
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		BirthDate: body.BirthDate,
		Password:  body.Password,
	}
	return v, nil
}

// BuildCustomerLoginPayload builds the payload for the coachee CustomerLogin
// endpoint from CLI flags.
func BuildCustomerLoginPayload(coacheeCustomerLoginBody string) (*coachee.CustomerLoginPayload, error) {
	var err error
	var body CustomerLoginRequestBody
	{
		err = json.Unmarshal([]byte(coacheeCustomerLoginBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"email\": \"In aut quia recusandae minima eius.\",\n      \"password\": \"Omnis deleniti praesentium at culpa.\"\n   }'")
		}
	}
	v := &coachee.CustomerLoginPayload{
		Email:    body.Email,
		Password: body.Password,
	}
	return v, nil
}

// BuildStartPasswordRecoveryFlowPayload builds the payload for the coachee
// StartPasswordRecoveryFlow endpoint from CLI flags.
func BuildStartPasswordRecoveryFlowPayload(coacheeStartPasswordRecoveryFlowBody string) (*coachee.StartPasswordRecoveryFlowPayload, error) {
	var err error
	var body StartPasswordRecoveryFlowRequestBody
	{
		err = json.Unmarshal([]byte(coacheeStartPasswordRecoveryFlowBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"email\": \"Enim voluptatem atque expedita.\"\n   }'")
		}
	}
	v := &coachee.StartPasswordRecoveryFlowPayload{
		Email: body.Email,
	}
	return v, nil
}

// BuildCheckPasswordRecoveryTokenPayload builds the payload for the coachee
// CheckPasswordRecoveryToken endpoint from CLI flags.
func BuildCheckPasswordRecoveryTokenPayload(coacheeCheckPasswordRecoveryTokenToken string) (*coachee.CheckPasswordRecoveryTokenPayload, error) {
	var token string
	{
		token = coacheeCheckPasswordRecoveryTokenToken
	}
	payload := &coachee.CheckPasswordRecoveryTokenPayload{
		Token: token,
	}
	return payload, nil
}

// BuildFinalizePasswordRecoveryFlowPayload builds the payload for the coachee
// FinalizePasswordRecoveryFlow endpoint from CLI flags.
func BuildFinalizePasswordRecoveryFlowPayload(coacheeFinalizePasswordRecoveryFlowBody string, coacheeFinalizePasswordRecoveryFlowToken string) (*coachee.FinalizePasswordRecoveryFlowPayload, error) {
	var err error
	var body FinalizePasswordRecoveryFlowRequestBody
	{
		err = json.Unmarshal([]byte(coacheeFinalizePasswordRecoveryFlowBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"password\": \"Est occaecati molestiae praesentium illum et rerum.\"\n   }'")
		}
	}
	var token string
	{
		token = coacheeFinalizePasswordRecoveryFlowToken
	}
	v := &coachee.FinalizePasswordRecoveryFlowPayload{
		Password: body.Password,
	}
	v.Token = token
	return v, nil
}

// BuildCreateOrderPayload builds the payload for the coachee CreateOrder
// endpoint from CLI flags.
func BuildCreateOrderPayload(coacheeCreateOrderBody string, coacheeCreateOrderToken string) (*coachee.CreateOrderPayload, error) {
	var err error
	var body CreateOrderRequestBody
	{
		err = json.Unmarshal([]byte(coacheeCreateOrderBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"coachId\": 5337483527812557295,\n      \"introCall\": 5877709250813620535,\n      \"programId\": \"Necessitatibus suscipit officia dolor.\"\n   }'")
		}
	}
	var token string
	{
		token = coacheeCreateOrderToken
	}
	v := &coachee.CreateOrderPayload{
		CoachID:   body.CoachID,
		ProgramID: body.ProgramID,
		IntroCall: body.IntroCall,
	}
	v.Token = token
	return v, nil
}

// BuildRegisterStripeExpressPayload builds the payload for the coachee
// RegisterStripeExpress endpoint from CLI flags.
func BuildRegisterStripeExpressPayload(coacheeRegisterStripeExpressBody string, coacheeRegisterStripeExpressID string) (*coachee.RegisterStripeExpressPayload, error) {
	var err error
	var body RegisterStripeExpressRequestBody
	{
		err = json.Unmarshal([]byte(coacheeRegisterStripeExpressBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"authorizationCode\": \"Et perspiciatis.\"\n   }'")
		}
	}
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(coacheeRegisterStripeExpressID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &coachee.RegisterStripeExpressPayload{
		AuthorizationCode: body.AuthorizationCode,
	}
	v.ID = id
	return v, nil
}

// BuildAdminLoginPayload builds the payload for the coachee AdminLogin
// endpoint from CLI flags.
func BuildAdminLoginPayload(coacheeAdminLoginBody string) (*coachee.AdminLoginPayload, error) {
	var err error
	var body AdminLoginRequestBody
	{
		err = json.Unmarshal([]byte(coacheeAdminLoginBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"email\": \"Eum ipsam.\",\n      \"password\": \"Illo saepe et nihil.\"\n   }'")
		}
	}
	v := &coachee.AdminLoginPayload{
		Email:    body.Email,
		Password: body.Password,
	}
	return v, nil
}
