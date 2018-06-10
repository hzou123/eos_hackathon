package router

import (
	"net/http"
	"../controller"
)

// Define a single route, e.g. a human readable name, HTTP method and the pattern the function that will execute when the route is called.

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

var UserRoutes = Routes{
	Route{
		"GetUser", // Name
		"GET",        // HTTP method
		"/getUserBalance/{uname}", // Route pattern
		controller.GetUserBalance,
	},

}

var ProductRoutes = Routes{
	Route{
		"GetQuoteInfo", // Name
		"GET",        // HTTP method
		"/GetQuoteInfo/{projectName}", // Route pattern
		controller.GetQuoteInfo,
	},
}

var OrderRoutes = Routes{
	Route{
		"submitOrder", // Name
		"GET",        // HTTP method
		"/submitOrder/{pid}", // Route pattern
		controller.MakeOrder,
	},
}

