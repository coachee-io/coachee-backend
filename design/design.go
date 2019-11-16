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
})

var certifications = Type("certifications", func() {
	Description("represents a coach certification")

	Attribute("id", String)
	Attribute("title", String)
	Attribute("description", String)
	Attribute("institution", String)
	Attribute("month", UInt, func() {
		Minimum(uint(1))
		Maximum(uint(12))
	})
	Attribute("Year", UInt, func() {
		Minimum(uint(1900))
		Maximum(uint(2100))
	})

	Required("title", "description", "institution", "month", "year")
})

//var hour = Type("hour", UInt, func() {
//	Minimum(uint(0))
//	Maximum(uint(24))
//})
//
//var minute = Type("name", UInt, func() {
//	Minimum(uint(0))
//	Maximum(uint(59))
//})

var availability = Type("availability", func() {
	Description("represents a coach availability")

	Attribute("id", String)
	Attribute("weekDay", String)
	Attribute("startHour", UInt)
	Attribute("endHour", UInt)
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

var coachResult = ResultType("coach", func() {
	Description("represents a coach and all his relevant info")
	ContentType("application/json")

	Attribute("id", UInt)
	Attribute("firstName", String)
	Attribute("lastName", String)
	Attribute("tags", String)
	Attribute("description", String)
	Attribute("city", String)
	Attribute("country", String)
	Attribute("pictureURL", String)
	Attribute("certifications", ArrayOf(certifications))
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

})
