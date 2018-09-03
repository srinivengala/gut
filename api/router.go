package api

import "github.com/ant0ine/go-json-rest/rest"

// NewRouter creates a rest router
func NewRouter() (rest.App, error) {

	var restRoutes []*rest.Route

	for _, route := range routesInstance {
		var handler rest.HandlerFunc
		//TODO add context
		handler = route.HandlerFunc

		restRoute := &rest.Route{
			HttpMethod: route.Method,
			PathExp:    route.Pattern,
			Func:       handler,
		}

		//restRoute := rest.Get(route.Pattern, handler)
		restRoutes = append(restRoutes, restRoute)
	}

	router, err := rest.MakeRouter(restRoutes...)

	// router, err := rest.MakeRouter(
	// 	//rest.Get("/users/:id", GetUser),
	// 	rest.Get("/hello", Hello),
	// )
	return router, err
}
