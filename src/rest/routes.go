package rest

import (
	"net/http"
	"github.com/andrepinto/navyhook/src/rest/controllers"
)

var routes = Routes{
	Route{
		"ApiVersion",
		"GET",
		"/api/version",
		controllers.ShowApiVersion,
	},
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route