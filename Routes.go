package main

import "net/http"

//Route is a type representing each route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//StaticRoute is a type representing a route to a folder
type StaticRoute struct {
	Name    string
	Method  string
	Pattern string
	Dir     string
}

//Routes is a type containing an array of type Route
type Routes []Route

//StaticRoutes is a type containing an array of type StaticRoute
type StaticRoutes []StaticRoute

var routes = Routes{
	Route{
		"NewGame",
		"GET",
		"/newgame",
		NewGame,
	},
}

var static = StaticRoutes{
	StaticRoute{
		"Index",
		"GET",
		"/",
		"static/",
	},
}
