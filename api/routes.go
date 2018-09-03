package api

// concrete rest api routes
var routesInstance = RoutesModel{
	RouteModel{"Index", "GET", "/", Hello},
	RouteModel{"Hello", "GET", "/hello", Hello},
	RouteModel{"Hello", "GET", "/hello/", Hello},
}
