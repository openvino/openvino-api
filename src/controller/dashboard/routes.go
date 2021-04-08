package dashboard

import (
	"net/http"

	customHTTP "github.com/openvino/openvino-api/src/http"
)

// Routes - Client related routes
var Routes = customHTTP.RoutePrefix{
	Prefix: "/dashboard",
	SubRoutes: []customHTTP.Route{
		{
			Name:        "GetDashboard",
			Method:      http.MethodGet,
			Pattern:     "",
			HandlerFunc: GetDashboard,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
	},
}
