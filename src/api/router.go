package api

import (
	"net/http"
	"github.com/gorilla/mux"
)

func ApiRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range GetRoutes() {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)

	}

	return router
}