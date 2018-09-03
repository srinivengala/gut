package api

import "github.com/ant0ine/go-json-rest/rest"

// RouteModel is Model for a route in router setup
type RouteModel struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc rest.HandlerFunc
}

// RoutesModel is Model for router setup
type RoutesModel []RouteModel
