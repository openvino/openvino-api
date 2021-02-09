package http

import (
	"net/http"
)

// AppRoutes - Global app routes
var AppRoutes []RoutePrefix

// RoutePrefix - Common route prefix type
type RoutePrefix struct {
	Prefix    string
	SubRoutes []Route
}

// Route - Common route type
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Scopes      []Scope
}
