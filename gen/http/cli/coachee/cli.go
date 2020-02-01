// Code generated by goa v3.0.9, DO NOT EDIT.
//
// coachee HTTP client CLI support package
//
// Command:
// $ goa gen coachee-backend/design

package cli

import (
	coacheec "coachee-backend/gen/http/coachee/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `coachee (get-coaches|get-coach|len-coaches|create-coach|update-coach|create-certification|delete-certification|create-program|delete-program|create-availability|delete-availability|create-customer|customer-login|start-password-recovery-flow|check-password-recovery-token|finalize-password-recovery-flow|create-order|register-stripe-express)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` coachee get-coaches --tag "Debitis molestiae et sit eos saepe." --limit 6604467783455990843 --page 17042997039165179651` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		coacheeFlags = flag.NewFlagSet("coachee", flag.ContinueOnError)

		coacheeGetCoachesFlags     = flag.NewFlagSet("get-coaches", flag.ExitOnError)
		coacheeGetCoachesTagFlag   = coacheeGetCoachesFlags.String("tag", "", "")
		coacheeGetCoachesLimitFlag = coacheeGetCoachesFlags.String("limit", "", "")
		coacheeGetCoachesPageFlag  = coacheeGetCoachesFlags.String("page", "", "")

		coacheeGetCoachFlags  = flag.NewFlagSet("get-coach", flag.ExitOnError)
		coacheeGetCoachIDFlag = coacheeGetCoachFlags.String("id", "REQUIRED", "")

		coacheeLenCoachesFlags   = flag.NewFlagSet("len-coaches", flag.ExitOnError)
		coacheeLenCoachesTagFlag = coacheeLenCoachesFlags.String("tag", "REQUIRED", "")

		coacheeCreateCoachFlags    = flag.NewFlagSet("create-coach", flag.ExitOnError)
		coacheeCreateCoachBodyFlag = coacheeCreateCoachFlags.String("body", "REQUIRED", "")

		coacheeUpdateCoachFlags    = flag.NewFlagSet("update-coach", flag.ExitOnError)
		coacheeUpdateCoachBodyFlag = coacheeUpdateCoachFlags.String("body", "REQUIRED", "")
		coacheeUpdateCoachIDFlag   = coacheeUpdateCoachFlags.String("id", "REQUIRED", "")

		coacheeCreateCertificationFlags    = flag.NewFlagSet("create-certification", flag.ExitOnError)
		coacheeCreateCertificationBodyFlag = coacheeCreateCertificationFlags.String("body", "REQUIRED", "")
		coacheeCreateCertificationIDFlag   = coacheeCreateCertificationFlags.String("id", "REQUIRED", "")

		coacheeDeleteCertificationFlags      = flag.NewFlagSet("delete-certification", flag.ExitOnError)
		coacheeDeleteCertificationIDFlag     = coacheeDeleteCertificationFlags.String("id", "REQUIRED", "")
		coacheeDeleteCertificationCertIDFlag = coacheeDeleteCertificationFlags.String("cert-id", "REQUIRED", "")

		coacheeCreateProgramFlags    = flag.NewFlagSet("create-program", flag.ExitOnError)
		coacheeCreateProgramBodyFlag = coacheeCreateProgramFlags.String("body", "REQUIRED", "")
		coacheeCreateProgramIDFlag   = coacheeCreateProgramFlags.String("id", "REQUIRED", "")

		coacheeDeleteProgramFlags         = flag.NewFlagSet("delete-program", flag.ExitOnError)
		coacheeDeleteProgramIDFlag        = coacheeDeleteProgramFlags.String("id", "REQUIRED", "")
		coacheeDeleteProgramProgramIDFlag = coacheeDeleteProgramFlags.String("program-id", "REQUIRED", "")

		coacheeCreateAvailabilityFlags    = flag.NewFlagSet("create-availability", flag.ExitOnError)
		coacheeCreateAvailabilityBodyFlag = coacheeCreateAvailabilityFlags.String("body", "REQUIRED", "")
		coacheeCreateAvailabilityIDFlag   = coacheeCreateAvailabilityFlags.String("id", "REQUIRED", "")

		coacheeDeleteAvailabilityFlags    = flag.NewFlagSet("delete-availability", flag.ExitOnError)
		coacheeDeleteAvailabilityIDFlag   = coacheeDeleteAvailabilityFlags.String("id", "REQUIRED", "")
		coacheeDeleteAvailabilityAvIDFlag = coacheeDeleteAvailabilityFlags.String("av-id", "REQUIRED", "")

		coacheeCreateCustomerFlags    = flag.NewFlagSet("create-customer", flag.ExitOnError)
		coacheeCreateCustomerBodyFlag = coacheeCreateCustomerFlags.String("body", "REQUIRED", "")

		coacheeCustomerLoginFlags    = flag.NewFlagSet("customer-login", flag.ExitOnError)
		coacheeCustomerLoginBodyFlag = coacheeCustomerLoginFlags.String("body", "REQUIRED", "")

		coacheeStartPasswordRecoveryFlowFlags    = flag.NewFlagSet("start-password-recovery-flow", flag.ExitOnError)
		coacheeStartPasswordRecoveryFlowBodyFlag = coacheeStartPasswordRecoveryFlowFlags.String("body", "REQUIRED", "")

		coacheeCheckPasswordRecoveryTokenFlags     = flag.NewFlagSet("check-password-recovery-token", flag.ExitOnError)
		coacheeCheckPasswordRecoveryTokenTokenFlag = coacheeCheckPasswordRecoveryTokenFlags.String("token", "REQUIRED", "")

		coacheeFinalizePasswordRecoveryFlowFlags     = flag.NewFlagSet("finalize-password-recovery-flow", flag.ExitOnError)
		coacheeFinalizePasswordRecoveryFlowBodyFlag  = coacheeFinalizePasswordRecoveryFlowFlags.String("body", "REQUIRED", "")
		coacheeFinalizePasswordRecoveryFlowTokenFlag = coacheeFinalizePasswordRecoveryFlowFlags.String("token", "REQUIRED", "")

		coacheeCreateOrderFlags     = flag.NewFlagSet("create-order", flag.ExitOnError)
		coacheeCreateOrderBodyFlag  = coacheeCreateOrderFlags.String("body", "REQUIRED", "")
		coacheeCreateOrderTokenFlag = coacheeCreateOrderFlags.String("token", "REQUIRED", "")

		coacheeRegisterStripeExpressFlags    = flag.NewFlagSet("register-stripe-express", flag.ExitOnError)
		coacheeRegisterStripeExpressBodyFlag = coacheeRegisterStripeExpressFlags.String("body", "REQUIRED", "")
		coacheeRegisterStripeExpressIDFlag   = coacheeRegisterStripeExpressFlags.String("id", "REQUIRED", "")
	)
	coacheeFlags.Usage = coacheeUsage
	coacheeGetCoachesFlags.Usage = coacheeGetCoachesUsage
	coacheeGetCoachFlags.Usage = coacheeGetCoachUsage
	coacheeLenCoachesFlags.Usage = coacheeLenCoachesUsage
	coacheeCreateCoachFlags.Usage = coacheeCreateCoachUsage
	coacheeUpdateCoachFlags.Usage = coacheeUpdateCoachUsage
	coacheeCreateCertificationFlags.Usage = coacheeCreateCertificationUsage
	coacheeDeleteCertificationFlags.Usage = coacheeDeleteCertificationUsage
	coacheeCreateProgramFlags.Usage = coacheeCreateProgramUsage
	coacheeDeleteProgramFlags.Usage = coacheeDeleteProgramUsage
	coacheeCreateAvailabilityFlags.Usage = coacheeCreateAvailabilityUsage
	coacheeDeleteAvailabilityFlags.Usage = coacheeDeleteAvailabilityUsage
	coacheeCreateCustomerFlags.Usage = coacheeCreateCustomerUsage
	coacheeCustomerLoginFlags.Usage = coacheeCustomerLoginUsage
	coacheeStartPasswordRecoveryFlowFlags.Usage = coacheeStartPasswordRecoveryFlowUsage
	coacheeCheckPasswordRecoveryTokenFlags.Usage = coacheeCheckPasswordRecoveryTokenUsage
	coacheeFinalizePasswordRecoveryFlowFlags.Usage = coacheeFinalizePasswordRecoveryFlowUsage
	coacheeCreateOrderFlags.Usage = coacheeCreateOrderUsage
	coacheeRegisterStripeExpressFlags.Usage = coacheeRegisterStripeExpressUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "coachee":
			svcf = coacheeFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "coachee":
			switch epn {
			case "get-coaches":
				epf = coacheeGetCoachesFlags

			case "get-coach":
				epf = coacheeGetCoachFlags

			case "len-coaches":
				epf = coacheeLenCoachesFlags

			case "create-coach":
				epf = coacheeCreateCoachFlags

			case "update-coach":
				epf = coacheeUpdateCoachFlags

			case "create-certification":
				epf = coacheeCreateCertificationFlags

			case "delete-certification":
				epf = coacheeDeleteCertificationFlags

			case "create-program":
				epf = coacheeCreateProgramFlags

			case "delete-program":
				epf = coacheeDeleteProgramFlags

			case "create-availability":
				epf = coacheeCreateAvailabilityFlags

			case "delete-availability":
				epf = coacheeDeleteAvailabilityFlags

			case "create-customer":
				epf = coacheeCreateCustomerFlags

			case "customer-login":
				epf = coacheeCustomerLoginFlags

			case "start-password-recovery-flow":
				epf = coacheeStartPasswordRecoveryFlowFlags

			case "check-password-recovery-token":
				epf = coacheeCheckPasswordRecoveryTokenFlags

			case "finalize-password-recovery-flow":
				epf = coacheeFinalizePasswordRecoveryFlowFlags

			case "create-order":
				epf = coacheeCreateOrderFlags

			case "register-stripe-express":
				epf = coacheeRegisterStripeExpressFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "coachee":
			c := coacheec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-coaches":
				endpoint = c.GetCoaches()
				data, err = coacheec.BuildGetCoachesPayload(*coacheeGetCoachesTagFlag, *coacheeGetCoachesLimitFlag, *coacheeGetCoachesPageFlag)
			case "get-coach":
				endpoint = c.GetCoach()
				data, err = coacheec.BuildGetCoachPayload(*coacheeGetCoachIDFlag)
			case "len-coaches":
				endpoint = c.LenCoaches()
				data, err = coacheec.BuildLenCoachesPayload(*coacheeLenCoachesTagFlag)
			case "create-coach":
				endpoint = c.CreateCoach()
				data, err = coacheec.BuildCreateCoachPayload(*coacheeCreateCoachBodyFlag)
			case "update-coach":
				endpoint = c.UpdateCoach()
				data, err = coacheec.BuildUpdateCoachPayload(*coacheeUpdateCoachBodyFlag, *coacheeUpdateCoachIDFlag)
			case "create-certification":
				endpoint = c.CreateCertification()
				data, err = coacheec.BuildCreateCertificationPayload(*coacheeCreateCertificationBodyFlag, *coacheeCreateCertificationIDFlag)
			case "delete-certification":
				endpoint = c.DeleteCertification()
				data, err = coacheec.BuildDeleteCertificationPayload(*coacheeDeleteCertificationIDFlag, *coacheeDeleteCertificationCertIDFlag)
			case "create-program":
				endpoint = c.CreateProgram()
				data, err = coacheec.BuildCreateProgramPayload(*coacheeCreateProgramBodyFlag, *coacheeCreateProgramIDFlag)
			case "delete-program":
				endpoint = c.DeleteProgram()
				data, err = coacheec.BuildDeleteProgramPayload(*coacheeDeleteProgramIDFlag, *coacheeDeleteProgramProgramIDFlag)
			case "create-availability":
				endpoint = c.CreateAvailability()
				data, err = coacheec.BuildCreateAvailabilityPayload(*coacheeCreateAvailabilityBodyFlag, *coacheeCreateAvailabilityIDFlag)
			case "delete-availability":
				endpoint = c.DeleteAvailability()
				data, err = coacheec.BuildDeleteAvailabilityPayload(*coacheeDeleteAvailabilityIDFlag, *coacheeDeleteAvailabilityAvIDFlag)
			case "create-customer":
				endpoint = c.CreateCustomer()
				data, err = coacheec.BuildCreateCustomerPayload(*coacheeCreateCustomerBodyFlag)
			case "customer-login":
				endpoint = c.CustomerLogin()
				data, err = coacheec.BuildCustomerLoginPayload(*coacheeCustomerLoginBodyFlag)
			case "start-password-recovery-flow":
				endpoint = c.StartPasswordRecoveryFlow()
				data, err = coacheec.BuildStartPasswordRecoveryFlowPayload(*coacheeStartPasswordRecoveryFlowBodyFlag)
			case "check-password-recovery-token":
				endpoint = c.CheckPasswordRecoveryToken()
				data, err = coacheec.BuildCheckPasswordRecoveryTokenPayload(*coacheeCheckPasswordRecoveryTokenTokenFlag)
			case "finalize-password-recovery-flow":
				endpoint = c.FinalizePasswordRecoveryFlow()
				data, err = coacheec.BuildFinalizePasswordRecoveryFlowPayload(*coacheeFinalizePasswordRecoveryFlowBodyFlag, *coacheeFinalizePasswordRecoveryFlowTokenFlag)
			case "create-order":
				endpoint = c.CreateOrder()
				data, err = coacheec.BuildCreateOrderPayload(*coacheeCreateOrderBodyFlag, *coacheeCreateOrderTokenFlag)
			case "register-stripe-express":
				endpoint = c.RegisterStripeExpress()
				data, err = coacheec.BuildRegisterStripeExpressPayload(*coacheeRegisterStripeExpressBodyFlag, *coacheeRegisterStripeExpressIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// coacheeUsage displays the usage of the coachee command and its subcommands.
func coacheeUsage() {
	fmt.Fprintf(os.Stderr, `The coachee service performs operations on coachees
Usage:
    %s [globalflags] coachee COMMAND [flags]

COMMAND:
    get-coaches: GetCoaches returns an array of coaches according to a tag and pagination
    get-coach: GetCoach returns one coach according to the id
    len-coaches: LenCoaches returns the amount of coaches with a given tag
    create-coach: CreateCoaches creates a base coach
    update-coach: UpdateCoaches updates a coach
    create-certification: creates a certification for a coach
    delete-certification: deletes a certification for a coach
    create-program: creates a program for a coach
    delete-program: deletes a program for a coach
    create-availability: creates an availability for a coach
    delete-availability: deletes an availability for a coach
    create-customer: creates a new customer
    customer-login: logs in a customer and returns a jwt
    start-password-recovery-flow: starts the process of recovering a password
    check-password-recovery-token: verifies if a recovery token is still valid
    finalize-password-recovery-flow: finalizes the password recovery flow by resetting a new password 
    create-order: creates a new order
    register-stripe-express: registers a stripe express account in stripe and associates it to a coach

Additional help:
    %s coachee COMMAND --help
`, os.Args[0], os.Args[0])
}
func coacheeGetCoachesUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee get-coaches -tag STRING -limit UINT -page UINT

GetCoaches returns an array of coaches according to a tag and pagination
    -tag STRING: 
    -limit UINT: 
    -page UINT: 

Example:
    `+os.Args[0]+` coachee get-coaches --tag "Debitis molestiae et sit eos saepe." --limit 6604467783455990843 --page 17042997039165179651
`, os.Args[0])
}

func coacheeGetCoachUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee get-coach -id UINT

GetCoach returns one coach according to the id
    -id UINT: 

Example:
    `+os.Args[0]+` coachee get-coach --id 11687188891161822163
`, os.Args[0])
}

func coacheeLenCoachesUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee len-coaches -tag STRING

LenCoaches returns the amount of coaches with a given tag
    -tag STRING: 

Example:
    `+os.Args[0]+` coachee len-coaches --tag "Iusto rerum voluptas rerum amet id deserunt."
`, os.Args[0])
}

func coacheeCreateCoachUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee create-coach -body JSON

CreateCoaches creates a base coach
    -body JSON: 

Example:
    `+os.Args[0]+` coachee create-coach --body '{
      "city": "Voluptates tenetur in ut qui blanditiis quidem.",
      "country": "Accusamus mollitia reiciendis tenetur rerum nulla optio.",
      "description": "Quidem sint.",
      "email": "Nam facilis aut veritatis aut.",
      "firstName": "Dolor sunt eos.",
      "introCall": 7033531370616248077,
      "lastName": "Quia quia blanditiis.",
      "password": "Est culpa.",
      "phone": "Quia sit est enim unde illo dolorem.",
      "tags": "Cum et magnam.",
      "textAvailability": "Nisi quam.",
      "textCertifications": "Harum facere ipsa voluptate.",
      "textPrograms": "Excepturi natus nesciunt.",
      "vat": "Provident porro doloremque repellendus nemo."
   }'
`, os.Args[0])
}

func coacheeUpdateCoachUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee update-coach -body JSON -id UINT

UpdateCoaches updates a coach
    -body JSON: 
    -id UINT: 

Example:
    `+os.Args[0]+` coachee update-coach --body '{
      "city": "Quae magnam ut vero autem magnam rerum.",
      "country": "Autem harum aut beatae.",
      "description": "Rerum laborum.",
      "email": "Iusto incidunt odio.",
      "firstName": "Aliquid est tempora iure id aspernatur.",
      "introCall": 14037145708316104300,
      "lastName": "Esse molestiae eum et et.",
      "phone": "Quo incidunt ad recusandae quam.",
      "pictureURL": "Sint facere est.",
      "stripeID": "Id aperiam fugit facere repellat distinctio architecto.",
      "tags": "Sint facilis.",
      "vat": "Rerum necessitatibus debitis praesentium accusamus."
   }' --id 5767410731598171931
`, os.Args[0])
}

func coacheeCreateCertificationUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee create-certification -body JSON -id UINT

creates a certification for a coach
    -body JSON: 
    -id UINT: 

Example:
    `+os.Args[0]+` coachee create-certification --body '{
      "certification": {
         "description": "Quia totam quia minima doloremque necessitatibus pariatur.",
         "id": "Recusandae ab sint.",
         "institution": "Iste sint.",
         "month": 9,
         "title": "Est qui laborum qui.",
         "year": 2001
      }
   }' --id 13719383368494616160
`, os.Args[0])
}

func coacheeDeleteCertificationUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee delete-certification -id UINT -cert-id STRING

deletes a certification for a coach
    -id UINT: 
    -cert-id STRING: 

Example:
    `+os.Args[0]+` coachee delete-certification --id 17712998355210063996 --cert-id "Minima repudiandae occaecati quis error."
`, os.Args[0])
}

func coacheeCreateProgramUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee create-program -body JSON -id UINT

creates a program for a coach
    -body JSON: 
    -id UINT: 

Example:
    `+os.Args[0]+` coachee create-program --body '{
      "program": {
         "description": "Numquam velit dolor sit dolor.",
         "duration": 14591392219494929222,
         "id": "Velit eveniet.",
         "name": "Aliquid saepe velit suscipit delectus fugit.",
         "sessions": 8578413014643017611,
         "taxPercent": 8855644995256618976,
         "totalPrice": 17901160786247009004
      }
   }' --id 13303027061162062993
`, os.Args[0])
}

func coacheeDeleteProgramUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee delete-program -id UINT -program-id STRING

deletes a program for a coach
    -id UINT: 
    -program-id STRING: 

Example:
    `+os.Args[0]+` coachee delete-program --id 6800996720000124411 --program-id "Quia non commodi."
`, os.Args[0])
}

func coacheeCreateAvailabilityUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee create-availability -body JSON -id UINT

creates an availability for a coach
    -body JSON: 
    -id UINT: 

Example:
    `+os.Args[0]+` coachee create-availability --body '{
      "availability": {
         "end": 465,
         "id": "Rerum officiis.",
         "start": 184,
         "weekDay": 0
      }
   }' --id 3913520442231866481
`, os.Args[0])
}

func coacheeDeleteAvailabilityUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee delete-availability -id UINT -av-id STRING

deletes an availability for a coach
    -id UINT: 
    -av-id STRING: 

Example:
    `+os.Args[0]+` coachee delete-availability --id 1600104616427973716 --av-id "Aut beatae non quia blanditiis similique."
`, os.Args[0])
}

func coacheeCreateCustomerUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee create-customer -body JSON

creates a new customer
    -body JSON: 

Example:
    `+os.Args[0]+` coachee create-customer --body '{
      "birthDate": 2951058717071938166,
      "email": "Placeat consequatur cum.",
      "firstName": "Quo id dolorum reprehenderit exercitationem est.",
      "lastName": "Harum eius.",
      "password": "Sapiente deserunt molestiae."
   }'
`, os.Args[0])
}

func coacheeCustomerLoginUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee customer-login -body JSON

logs in a customer and returns a jwt
    -body JSON: 

Example:
    `+os.Args[0]+` coachee customer-login --body '{
      "email": "Tempora modi.",
      "password": "Quos eum rerum architecto et facere."
   }'
`, os.Args[0])
}

func coacheeStartPasswordRecoveryFlowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee start-password-recovery-flow -body JSON

starts the process of recovering a password
    -body JSON: 

Example:
    `+os.Args[0]+` coachee start-password-recovery-flow --body '{
      "email": "Nihil tempora."
   }'
`, os.Args[0])
}

func coacheeCheckPasswordRecoveryTokenUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee check-password-recovery-token -token STRING

verifies if a recovery token is still valid
    -token STRING: 

Example:
    `+os.Args[0]+` coachee check-password-recovery-token --token "Sunt doloremque veritatis vel."
`, os.Args[0])
}

func coacheeFinalizePasswordRecoveryFlowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee finalize-password-recovery-flow -body JSON -token STRING

finalizes the password recovery flow by resetting a new password 
    -body JSON: 
    -token STRING: 

Example:
    `+os.Args[0]+` coachee finalize-password-recovery-flow --body '{
      "password": "Harum harum vitae iusto illo sit voluptatem."
   }' --token "Qui omnis consequatur in omnis et ex."
`, os.Args[0])
}

func coacheeCreateOrderUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee create-order -body JSON -token STRING

creates a new order
    -body JSON: 
    -token STRING: 

Example:
    `+os.Args[0]+` coachee create-order --body '{
      "coachId": 6707329220519220526,
      "introCall": 4123726178289067981,
      "programId": "Aut optio sunt blanditiis illum in."
   }' --token "Vel eligendi magnam repudiandae consequuntur explicabo."
`, os.Args[0])
}

func coacheeRegisterStripeExpressUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] coachee register-stripe-express -body JSON -id UINT

registers a stripe express account in stripe and associates it to a coach
    -body JSON: 
    -id UINT: 

Example:
    `+os.Args[0]+` coachee register-stripe-express --body '{
      "expressId": "Amet reiciendis voluptates et sed harum."
   }' --id 18252683255504708559
`, os.Args[0])
}
