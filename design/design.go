package design

import . "goa.design/goa/v3/dsl"

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

var certification = Type("certifications", func() {
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
})

var _ = Service("coachee", func() {
	Description("The coachee service performs operations on coachees")

	Method("GetCoaches", func() {
		Description("GetCoaches returns an array of coaches according to a tag and pagination")
		Payload(func() {
			Attribute("tag", String)
			Attribute("limit", UInt)
			Attribute("page", UInt)
			Required("tag")
		})

		Result(coachResult)

		HTTP(func() {
			GET("/coaches/{tag}")
			Param("limit")
			Param("page")
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
			Attribute("certifications", ArrayOf(certification))
			Attribute("programs", ArrayOf(program))
			Attribute("introCall", UInt) // maybe an external scheduler

			Required("firstName", "lastName", "email", "phone", "tags", "description",
				"certifications", "programs", "introCall")
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
			PUT("/coaches/{id}/certifications")
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
			PUT("/coaches/{id}/programs")
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
			PUT("/coaches/{id}/availability")
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
})
