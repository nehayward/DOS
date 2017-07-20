package api

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"AttackList",
		http.MethodGet,
		"/attacks",
		AttackList,
	},
	Route{
		"AttackShow",
		http.MethodGet,
		"/attacks/{attackid}",
		AttackShow,
	},
	Route{
		"AttackCreate",
		http.MethodPost,
		"/attack",
		AttackCreate,
	},
	Route{
		"AttackStop",
		http.MethodDelete,
		"/stop",
		AttackStop,
	},
}
