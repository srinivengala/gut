package api

import "github.com/ant0ine/go-json-rest/rest"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc rest.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"v1/",
		Hello,
	},
	Route{
		"Index",
		"GET",
		"v1/hello",
		Hello,
	},
}
