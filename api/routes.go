package api

import "github.com/srinivengala/gut/model"

// NewRoutes creates Routes Model instance
func (a *API) Routes() *model.Routes {
	return &model.Routes{
		model.Route{Name: "Index", Method: "GET", Pattern: "/", HandlerFunc: a.Hello},
		model.Route{Name: "Hello", Method: "GET", Pattern: "/hello", HandlerFunc: a.Hello},
		model.Route{Name: "Hello", Method: "GET", Pattern: "/hello/", HandlerFunc: a.Hello},
	}
}
