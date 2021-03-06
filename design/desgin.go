package design

import (
	. "goa.design/goa/dsl"
)

var _ = API("snake_case_fields", func() {
	Title("An example of snake case fields")

	Server("examples", func() {
		Services("user")

		// List the Hosts and their transport URLs.
		Host("development", func() {
			Description("Development hosts.")
			URI("http://localhost:8000/")
		})
	})
})

var fieldNames = []string{"first_name", "last_name"}

var UserPayload = Type("UserPayload", func() {
	for _, f := range fieldNames {
		Attribute(f, String)
	}
	Required("first_name")
})

var User = ResultType("application/vnd.user+json", func() {
	Attributes(func() {
		Attribute("id")
		for _, f := range fieldNames {
			Attribute(f, String)
		}
		Required(fieldNames...)
	})
	View("default", func() {
		Attribute("id")
		for _, f := range fieldNames {
			Attribute(f)
		}
	})
})

var _ = Service("user", func() {
	HTTP(func() {
		Path("/users")
	})

	Method("create", func() {
		Payload(UserPayload)
		Result(User)
		HTTP(func() {
			POST("")
			Body(UserPayload)
			Response(StatusCreated)
		})
	})

	Method("update", func() {
		Payload(func() {
			Attribute("id", Int)
			Attribute("user", UserPayload)
			Required("id", "user")
		})
		Result(User)
		HTTP(func() {
			PUT("/{id}")
			Body("user")
			Response(StatusOK)
		})
	})
})
