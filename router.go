package main

import (
	"net/http"

	"github.com/openvino/openvino-api/src/controller/dashboard"
	"github.com/openvino/openvino-api/src/controller/expense"
	"github.com/openvino/openvino-api/src/controller/language"
	"github.com/openvino/openvino-api/src/controller/redeem"
	"github.com/openvino/openvino-api/src/controller/sale"
	"github.com/openvino/openvino-api/src/controller/sensor"
	"github.com/openvino/openvino-api/src/controller/task"
	"github.com/openvino/openvino-api/src/controller/token"
	"github.com/openvino/openvino-api/src/controller/winerie"

	"github.com/openvino/openvino-api/src/controller/auth"
	"github.com/openvino/openvino-api/src/controller/health"
	customHTTP "github.com/openvino/openvino-api/src/http"
	"github.com/openvino/openvino-api/src/middleware"

	"github.com/gorilla/mux"
)

// NewRouter - Sets up a new router
func NewRouter() *mux.Router {

	router := mux.NewRouter()
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(s)

	customHTTP.AppRoutes = append(customHTTP.AppRoutes, health.Routes, auth.Routes, language.Routes, sale.Routes, sensor.Routes, task.Routes, redeem.Routes, dashboard.Routes, expense.Routes, winerie.Routes, token.Routes)

	for _, route := range customHTTP.AppRoutes {

		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		for _, r := range route.SubRoutes {

			var handler http.Handler
			handler = r.HandlerFunc

			handler = middleware.ContentTypeMiddleware(handler)
			handler = middleware.AuthMiddleware(handler, r.Scopes)
			handler = middleware.CorsMiddleware(handler)

			routePrefix.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method, "OPTIONS").
				Name(r.Name)
		}
	}

	return router
}
