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
			"ApiRoutes",
			"POST",
			"/api/configuration",
			api.CreateConfiguration,
		},
		Route{
			"ApiRoutes",
			"PUT",
			"/api/configuration/{name}",
			api.UpdateConfiguration,
		},
		Route{
			"ApiRoutes",
			"DELETE",
			"/api/configuration/{name}",
			api.DeleteConfiguration,
		},
		Route{
			"ApiRoutes",
			"POST",
			"/api/repos",
			api.AddRepo,
		},
		Route{
			"ApiRoutes",
			"DELETE",
			"/api/repos/{name}",
			api.DeleteRepo,
		},
		Route{
			"ApiRoutes",
			"GET",
			"/api/repos/{name}",
			api.GetRepo,
		},
		Route{
			"ApiRoutes",
			"GET",
			"/api/repos",
			api.GetAllRepos,
		},
		Route{
			"ApiRoutes",
			"POST",
			"/api/repos/{name}/hooks",
			api.CreateRepoHook,
		},
		Route{
			"ApiRoutes2",
			"PUT",
			"/api/repos/{name}/hooks/{id}",
			api.UpdateRepoHook,
		},
		Route{
			"ApiHooks",
			"POST",
			"/api/hook",
			api.HookReceive,
		},
		Route{
			"ApiHooks",
			"DELETE",
			"/api/repos/{name}/hooks/{id}",
			api.DeleteRepoHook,
		},
		Route{
			"ApiHooks",
			"GET",
			"/api/actions/repos/{repo}",
			api.GetRepoActions,
		},
		Route{
			"ApiHooks",
			"GET",
			"/api/actions/repos",
			api.GetAllRepoActions,
		},
	}

	return routes
}

