package api

import "github.com/ant0ine/go-json-rest/rest"

// NewRouter creates a rest router
func NewRouter() (rest.App, error) {
	router, err := rest.MakeRouter(
		//rest.Get("/users/:id", GetUser),
		rest.Get("/hello", Hello),
	)
	return router, err
}
