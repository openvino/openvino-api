package auth

import (
	"net/http"

	customHTTP "github.com/openvino/openvino-api/src/http"
)

// Routes - Client related routes
var Routes = customHTTP.RoutePrefix{
	Prefix: "/auth",
	SubRoutes: []customHTTP.Route{
		{
			Name:        "Auth",
			Method:      http.MethodGet,
			Pattern:     "",
			HandlerFunc: AuthHandler,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
	},
}
