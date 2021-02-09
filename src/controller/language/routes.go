package language

import (
	"net/http"

	customHTTP "github.com/openvino/openvino-api/src/http"
)

// Routes - Client related routes
var Routes = customHTTP.RoutePrefix{
	Prefix: "/language",
	SubRoutes: []customHTTP.Route{
		{
			Name:        "GetLanguage",
			Method:      http.MethodGet,
			Pattern:     "/{lang}",
			HandlerFunc: GetLanguage,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
	},
}
