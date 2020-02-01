package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

// API describes the global properties of the API server.
var _ = API("coachee", func() {
	Title("coachee service")
	Description("HTTP service for handling coaches, coachees and their transactions")
	Version("1.0")

	Server("coachee", func() {

		Services("coachee")

		Host("development", func() {
			Description("Development hosts.")
			URI("http://localhost:80")
			URI("grpc://localhost:8080")
		})
	})

	HTTP(func() {
		Consumes("application/json")
		Produces("application/json")
	})
})

var certification = Type("certification", func() {
	Description("represents a coach certification")

	Attribute("id", String)
	Attribute("title", String)
	Attribute("description", String)
	Attribute("institution", String)
	Attribute("month", UInt, func() {
		Minimum(1)
		Maximum(12)
	})
	Attribute("year", UInt, func() {
		Minimum(1900)
		Maximum(2100)
	})

	Required("title", "description", "institution", "month", "year")
})

var availability = Type("availability", func() {
	Description("represents a coach availability")

	Attribute("id", String)
	Attribute("weekDay", UInt, func() {
		Minimum(0)
		Maximum(6)
	})
	Attribute("start", UInt, func() {
		Minimum(0)
		Maximum(1440)
	})
	Attribute("end", UInt, func() {
		Minimum(0)
		Maximum(1440)
	})

	Required("weekDay", "start", "end")
})

var program = Type("program", func() {
	Description("represents a coach's programs")

	Attribute("id", String)
	Attribute("name", String)
	Attribute("sessions", UInt)
	Attribute("duration", UInt)
	Attribute("description", String)
	Attribute("totalPrice", UInt)
	Attribute("taxPercent", UInt)

	Required("name", "sessions", "duration", "description", "totalPrice", "taxPercent")
})

var coachResult = Type("coach", func() {
	Description("represents a coach and all his relevant info")

	Attribute("id", UInt)
	Attribute("firstName", String)
	Attribute("lastName", String)
	Attribute("tags", String)
	Attribute("description", String)
	Attribute("city", String)
	Attribute("country", String)
	Attribute("pictureURL", String)
	Attribute("certifications", ArrayOf(certification))
	Attribute("programs", ArrayOf(program))
	Attribute("availability", ArrayOf(availability))

	Required("id", "firstName", "lastName", "tags", "description", "city", "country", "pictureURL")
})

var customer = Type("baseClient", func() {
	Description("represents a client")

	Attribute("id", UInt)
	Attribute("firstName", String)
	Attribute("lastName", String)

	Required("id", "firstName", "lastName")
})

var JWT = JWTSecurity("jwt", func() {
	Scope("client", "client auth")
	Scope("admin", "admin auth")
})

var _ = Service("coachee", func() {
	cors.Origin("*", func() {
		cors.Headers("*")
		cors.Methods("GET", "POST", "PUT", "DELETE")
		cors.Credentials()
	})
	cors.Origin("localhost", func() {
		cors.Headers("*")
		cors.Methods("GET", "POST", "PUT", "DELETE")
		cors.Credentials()
	})

	Description("The coachee service performs operations on coachees")

	// error definition
	Error("transient", func() {
		Temporary()
	})
	Error("notFound")
	Error("validation")
	Error("unauthorized")
	Error("internal")
	HTTP(func() {
		Response("internal", StatusInternalServerError)
		Response("transient", StatusInternalServerError)
		Response("notFound", StatusNotFound)
		Response("validation", StatusBadRequest)
		Response("unauthorized", StatusUnauthorized)
	})

	Method("GetCoaches", func() {
		Description("GetCoaches returns an array of coaches according to a tag and pagination")
		Payload(func() {
			Attribute("tag", String)
			Attribute("limit", UInt)
			Attribute("page", UInt)
		})

		Result(ArrayOf(coachResult))

		HTTP(func() {
			GET("/coaches")
			Param("tag")
			Param("limit")
			Param("page")
			Response(StatusOK)
		})
	})

	Method("GetCoach", func() {
		Description("GetCoach returns one coach according to the id")
		Payload(func() {
			Attribute("id", UInt)
			Required("id")
		})

		Result(coachResult)

		HTTP(func() {
			GET("/coaches/{id}")
			Response(StatusOK)
		})
	})

	Method("LenCoaches", func() {
		Description("LenCoaches returns the amount of coaches with a given tag")
		Payload(func() {
			Attribute("tag", String)
			Required("tag")
		})

		Result(UInt)

		HTTP(func() {
			GET("/coaches/{tag}/length")
			Response(StatusOK)
		})
	})

	Method("CreateCoach", func() {
		Description("CreateCoaches creates a base coach")
		Payload(func() {
			Attribute("firstName", String)
			Attribute("lastName", String)
			Attribute("email", String)
			Attribute("phone", String)
			Attribute("tags", String)
			Attribute("description", String)
			Attribute("city", String)
			Attribute("country", String)
			Attribute("introCall", UInt) // maybe an external scheduler
			Attribute("textCertifications", String)
			Attribute("textPrograms", String)
			Attribute("textAvailability", String)
			Attribute("vat", String)

			Required("firstName", "lastName", "email", "phone", "tags", "description", "introCall",
				"textCertifications", "textPrograms")
		})

		Result(UInt)

		HTTP(func() {
			POST("/coaches")
			Response(StatusCreated)
		})
	})

	Method("UpdateCoach", func() {
		Description("UpdateCoaches updates a coach")
		Payload(func() {
			Attribute("id", UInt)
			Attribute("firstName", String)
			Attribute("lastName", String)
			Attribute("email", String)
			Attribute("phone", String)
			Attribute("tags", String)
			Attribute("description", String)
			Attribute("city", String)
			Attribute("country", String)
			Attribute("introCall", UInt)
			Attribute("stripeID", String)
			Attribute("pictureURL", String)
			Attribute("vat", String)

			Required("id")
		})

		HTTP(func() {
			POST("/coaches/{id}")
			Response(StatusAccepted)
		})
	})

	Method("CreateCertification", func() {
		Description("creates a certification for a coach")
		Payload(func() {
			Attribute("id", UInt)
			Attribute("certification", certification)

			Required("id", "certification")
		})

		HTTP(func() {
			POST("/coaches/{id}/certifications")
			Response(StatusAccepted)
		})
	})

	Method("DeleteCertification", func() {
		Description("deletes a certification for a coach")
		Payload(func() {
			Attribute("id", UInt)
			Attribute("certID", String)

			Required("id", "certID")
		})

		HTTP(func() {
			DELETE("/coaches/{id}/certifications/{certID}")
			Response(StatusOK)
		})
	})

	Method("CreateProgram", func() {
		Description("creates a program for a coach")
		Payload(func() {
			Attribute("id", UInt)
			Attribute("program", program)

			Required("id", "program")
		})

		HTTP(func() {
			POST("/coaches/{id}/programs")
			Response(StatusAccepted)
		})
	})

	Method("DeleteProgram", func() {
		Description("deletes a program for a coach")
		Payload(func() {
			Attribute("id", UInt)
			Attribute("programID", String)

			Required("id", "programID")
		})

		HTTP(func() {
			DELETE("/coaches/{id}/programs/{programID}")
			Response(StatusOK)
		})
	})

	Method("CreateAvailability", func() {
		Description("creates an availability for a coach")
		Payload(func() {
			Attribute("id", UInt)
			Attribute("availability", availability)

			Required("id", "availability")
		})

		HTTP(func() {
			POST("/coaches/{id}/availability")
			Response(StatusAccepted)
		})
	})

	Method("DeleteAvailability", func() {
		Description("deletes an availability for a coach")
		Payload(func() {
			Attribute("id", UInt)
			Attribute("avID", String)

			Required("id", "avID")
		})

		HTTP(func() {
			DELETE("/coaches/{id}/availability/{avID}")
			Response(StatusOK)
		})
	})

	Method("CreateCustomer", func() {
		Description("creates a new customer")
		Payload(func() {
			Attribute("email", String)
			Attribute("firstName", String)
			Attribute("lastName", String)
			Attribute("birthDate", Int64)
			Attribute("password", String)

			Required("email", "firstName", "lastName", "birthDate", "password")
		})

		Result(func() {
			Attribute("token", String)
			Attribute("expiry", Int64)
			Attribute("user", customer)

			Required("token", "expiry", "user")
		})

		HTTP(func() {
			POST("/clients")
			Response(StatusCreated)
		})
	})

	Method("CustomerLogin", func() {
		Description("logs in a customer and returns a jwt")
		Payload(func() {
			Attribute("email", String)
			Attribute("password", String)

			Required("email", "password")
		})

		Result(func() {
			Attribute("token", String)
			Attribute("expiry", Int64)
			Attribute("user", customer)

			Required("token", "expiry", "user")
		})

		HTTP(func() {
			POST("/clients/login")
			Response(StatusOK)
		})
	})

	Method("StartPasswordRecoveryFlow", func() {
		Description("starts the process of recovering a password")
		Payload(func() {
			Attribute("email", String)

			Required("email")
		})

		HTTP(func() {
			POST("/recovery")
			Response(StatusOK)
		})
	})

	Method("CheckPasswordRecoveryToken", func() {
		Description("verifies if a recovery token is still valid")
		Payload(func() {
			Attribute("token", String)

			Required("token")
		})

		HTTP(func() {
			GET("/recovery/{token}")
			Response(StatusOK)
		})
	})

	Method("FinalizePasswordRecoveryFlow", func() {
		Description("finalizes the password recovery flow by resetting a new password ")
		Payload(func() {
			Attribute("token", String)
			Attribute("password", String)

			Required("token", "password")
		})

		HTTP(func() {
			POST("/recovery/{token}")
			Response(StatusOK)
		})
	})

	Method("CreateOrder", func() {
		Description("creates a new order")
		Security(JWT, func() {
			Scope("client")
		})

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("coachId", UInt)
			Attribute("programId", String)
			Attribute("introCall", Int64)

			Required("token", "coachId", "programId", "introCall")
		})

		Result(func() {
			Attribute("clientSecret", String)
			Attribute("publishingKey", String)

			Required("clientSecret", "publishingKey")
		})

		HTTP(func() {
			POST("/orders")
			Response(StatusCreated)
		})
	})

	Method("RegisterStripeExpress", func() {
		Description("registers a stripe express account in stripe and associates it to a coach")

		Payload(func() {
			Attribute("id", UInt)
			Attribute("expressId", String)

			Required("id", "expressId")
		})

		HTTP(func() {
			POST("/coaches/{id}/stripe")
			Response(StatusCreated)
		})
	})
})
