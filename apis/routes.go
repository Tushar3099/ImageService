package apis

import (
	"net/http"

	"github.com/RetailPulse/services"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Post",
		"/api/submit", // makes a user in the database
		services.PostJob,
	},
	Route{
		"Get",
		"/api/status", // makes a user in the database
		services.GetJob,
	},
}
