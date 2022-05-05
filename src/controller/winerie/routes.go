package winerie

import (
	customHTTP "github.com/openvino/openvino-api/src/http"
	"net/http"
)

// Routes - Client related routes
var Routes = customHTTP.RoutePrefix{
	Prefix: "/wineries",
	SubRoutes: []customHTTP.Route{
		{
			Name:        "RegisterWinerie",
			Method:      http.MethodPost,
			Pattern:     "",
			HandlerFunc: CreateWinerie,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
		{
			Name:        "GetWineries",
			Method:      http.MethodGet,
			Pattern:     "",
			HandlerFunc: GetWineries,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
	},
}
