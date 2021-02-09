package health

import (
	"net/http"

	customHTTP "github.com/openvino/openvino-api/src/http"
)

// Routes - Client related routes
var Routes = customHTTP.RoutePrefix{
	Prefix: "/health",
	SubRoutes: []customHTTP.Route{
		{
			Name:        "HealthCheck",
			Method:      http.MethodGet,
			Pattern:     "",
			HandlerFunc: Handler,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
	},
}
