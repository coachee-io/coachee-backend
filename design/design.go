package design

import (
	. "goa.design/goa/v3/dsl"
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
	Attribute("start", Float64)
	Attribute("end", Float64)
	Attribute("dateLabel", String)

	Required("id", "weekDay", "start", "end", "dateLabel")
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

var fullCoach = Type("fullCoach", func() {
	Attribute("id", UInt)
	Attribute("firstName", String)
	Attribute("lastName", String)
	Attribute("email", String)
	Attribute("phone", String)
	Attribute("stripeID", String)
	Attribute("tags", String)
	Attribute("description", String)
	Attribute("city", String)
	Attribute("country", String)
	Attribute("pictureURL", String)
	Attribute("status", String)
	Attribute("vat", String)
	Attribute("introCall", Int)

	Attribute("availability", ArrayOf(availability))
	Attribute("certifications", ArrayOf(certification))
	Attribute("programs", ArrayOf(program))

	Required("id", "firstName", "lastName", "email", "phone", "stripeID", "tags", "description", "city",
		"country", "pictureURL", "status", "vat", "introCall", "availability", "certifications", "programs")
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

	Method("StripeWebhooks", func() {
		Description("Stripe webhook endpoint")
		Payload(Bytes)

		HTTP(func() {
			POST("/webhooks")
			Response(StatusOK)
		})
	})

	Method("GetCoaches", func() {
		Description("GetCoaches returns an array of coaches according to a tag and pagination")
		Payload(func() {
			Attribute("tag", String)
			Attribute("limit", UInt)
			Attribute("page", UInt)
			Attribute("show_all", Boolean)
		})

		Result(ArrayOf(coachResult))

		HTTP(func() {
			GET("/coaches")
			Param("tag")
			Param("limit")
			Param("page")
			Param("show_all")
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

	Method("AdminGetCoach", func() {
		Description("AdminGetCoach returns all coach info according to the id")
		Security(JWT, func() {
			Scope("admin")
		})

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("id", UInt)

			Required("id", "token")
		})

		Result(fullCoach)

		HTTP(func() {
			GET("/admin/coaches/{id}")
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
			Attribute("password", String)
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

			Required("firstName", "lastName", "email", "password", "phone", "tags", "description", "introCall",
				"textCertifications", "textPrograms")
		})

		Result(UInt)

		HTTP(func() {
			POST("/coaches")
			Response(StatusCreated)
		})
	})

	Method("LoginCoach", func() {
		Description("Logs in a coach to stripe express")
		Payload(func() {
			Attribute("email", String)
			Attribute("password", String)

			Required("email", "password")
		})

		Result(func() {
			Attribute("url", String)

			Required("url")
		})

		HTTP(func() {
			POST("/coaches/login")
			Response(StatusOK)
		})
	})

	Method("StartCoachPasswordRecoveryFlow", func() {
		Description("starts the process of recovering a password")
		Payload(func() {
			Attribute("email", String)

			Required("email")
		})

		HTTP(func() {
			POST("/coaches/recovery")
			Response(StatusOK)
		})
	})

	Method("CheckCoachPasswordRecoveryToken", func() {
		Description("verifies if a recovery token is still valid")
		Payload(func() {
			Attribute("token", String)

			Required("token")
		})

		HTTP(func() {
			GET("/coaches/recovery/{token}")
			Response(StatusOK)
		})
	})

	Method("FinalizeCoachPasswordRecoveryFlow", func() {
		Description("finalizes the password recovery flow by resetting a new password ")
		Payload(func() {
			Attribute("token", String)
			Attribute("password", String)

			Required("token", "password")
		})

		HTTP(func() {
			POST("/coaches/recovery/{token}")
			Response(StatusOK)
		})
	})

	Method("UpdateCoach", func() {
		Description("UpdateCoaches updates a coach")
		Security(JWT, func() {
			Scope("admin")
		})

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
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
			Attribute("status", String)

			Required("token", "id")
		})

		HTTP(func() {
			PUT("/admin/coaches/{id}")
			Response(StatusAccepted)
		})
	})

	Method("CreateCertification", func() {
		Description("creates a certification for a coach")
		Security(JWT, func() {
			Scope("admin")
		})

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("id", UInt)
			Attribute("certification", certification)

			Required("token", "id", "certification")
		})

		HTTP(func() {
			POST("/admin/coaches/{id}/certifications")
			Response(StatusAccepted)
		})
	})

	Method("DeleteCertification", func() {
		Description("deletes a certification for a coach")
		Security(JWT, func() {
			Scope("admin")
		})

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("id", UInt)
			Attribute("certID", String)

			Required("token", "id", "certID")
		})

		HTTP(func() {
			DELETE("/admin/coaches/{id}/certifications/{certID}")
			Response(StatusOK)
		})
	})

	Method("CreateProgram", func() {
		Description("creates a program for a coach")
		Security(JWT, func() {
			Scope("admin")
		})

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("id", UInt)
			Attribute("program", program)

			Required("token", "id", "program")
		})

		HTTP(func() {
			POST("/admin/coaches/{id}/programs")
			Response(StatusAccepted)
		})
	})

	Method("DeleteProgram", func() {
		Description("deletes a program for a coach")
		Security(JWT, func() {
			Scope("admin")
		})

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("id", UInt)
			Attribute("programID", String)

			Required("token", "id", "programID")
		})

		HTTP(func() {
			DELETE("/admin/coaches/{id}/programs/{programID}")
			Response(StatusOK)
		})
	})

	Method("CreateAvailability", func() {
		Description("creates an availability for a coach")
		Security(JWT, func() {
			Scope("admin")
		})

		Payload(func() {
			Description("represents a coach availability")
			Token("token", String, "JWT token used to perform authorization")
			Attribute("id", UInt)
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

			Required("id", "token", "weekDay", "start", "end")
		})

		HTTP(func() {
			POST("/admin/coaches/{id}/availability")
			Response(StatusAccepted)
		})
	})

	Method("DeleteAvailability", func() {
		Description("deletes an availability for a coach")
		Security(JWT, func() {
			Scope("admin")
		})

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("id", UInt)
			Attribute("avID", String)

			Required("token", "id", "avID")
		})

		HTTP(func() {
			DELETE("/admin/coaches/{id}/availability/{avID}")
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
			Attribute("authorizationCode", String)

			Required("id", "authorizationCode")
		})

		HTTP(func() {
			POST("/coaches/{id}/stripe")
			Response(StatusCreated)
		})
	})

	Method("AdminLogin", func() {
		Description("logs in a customer and returns a jwt")
		Payload(func() {
			Attribute("email", String)
			Attribute("password", String)

			Required("email", "password")
		})

		Result(func() {
			Attribute("token", String)
			Attribute("expiry", Int64)

			Required("token", "expiry")
		})

		HTTP(func() {
			POST("/admin/login")
			Response(StatusOK)
		})
	})
})
