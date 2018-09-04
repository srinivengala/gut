package model

import "github.com/ant0ine/go-json-rest/rest"

// Route is Model for a route in router setup
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc rest.HandlerFunc
}

// Routes is Model for router setup
type Routes []Route
