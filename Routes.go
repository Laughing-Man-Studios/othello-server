package main

import "net/http"

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type staticRoute struct {
	Name    string
	Method  string
	Pattern string
	Dir     string
}

type routes []route

type staticRoutes []staticRoute

var routeMap = routes{
	route{
		"NewGame",
		"GET",
		"/newgame",
		newGame,
	},
	route{
		"Move",
		"POST",
		"/move/{player}",
		move,
	},
	route{
		"Events",
		"GET",
		"/events",
		events,
	},
  route{
		"SetBoard",
		"Post",
		"/setboard",
		setboard,
	},
}

var static = staticRoutes{
	staticRoute{
		"Index",
		"GET",
		"/",
		"static/",
	},
}
