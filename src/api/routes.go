package api

import (

	"github.com/andrepinto/navyhook/src/api/controllers"
)

func GetRoutes() Routes{
	var routes = Routes{
		Route{
			"ApiVersion",
			"GET",
			"/api/version",
			api.ShowApiVersion,
		},
		Route{
			"ApiRoutes",
			"GET",
			"/api/config/repos",
			api.ShowConfigRepos,
		},
		Route{
			"ApiHooks",
			"POST",
			"/api/hook",
			api.HookReceive,
		},
	}

	return routes
}

