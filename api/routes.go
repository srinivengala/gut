package api

import "github.com/srinivengala/gut/model"

// NewRoutes creates Routes Model instance
func NewRoutes(api *API) *model.Routes {
	return &model.Routes{
		model.Route{Name: "Index", Method: "GET", Pattern: "/", HandlerFunc: api.Hello},
		model.Route{Name: "Hello", Method: "GET", Pattern: "/hello", HandlerFunc: api.Hello},
		model.Route{Name: "Hello", Method: "GET", Pattern: "/hello/", HandlerFunc: api.Hello},
	}
}

// // concrete rest api routes
// var routesInstance = model.Routes{
// 	model.Route{Name: "Index", Method: "GET", Pattern: "/", HandlerFunc: Hello},
// 	model.Route{Name: "Hello", Method: "GET", Pattern: "/hello", HandlerFunc: Hello},
// 	model.Route{Name: "Hello", Method: "GET", Pattern: "/hello/", HandlerFunc: Hello},
// }
