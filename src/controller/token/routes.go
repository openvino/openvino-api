package token

import (
	customHTTP "github.com/openvino/openvino-api/src/http"
	"net/http"
)

// Routes - Client related routes
var Routes = customHTTP.RoutePrefix{
	Prefix: "/token",
	SubRoutes: []customHTTP.Route{
		{
			Name:        "GetTokensByWinerie",
			Method:      http.MethodGet,
			Pattern:     "",
			HandlerFunc: GetTokensByWinerie,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
	},
}
