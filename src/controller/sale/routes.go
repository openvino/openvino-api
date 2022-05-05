package sale

import (
	"net/http"

	customHTTP "github.com/openvino/openvino-api/src/http"
)

// Routes - Client related routes
var Routes = customHTTP.RoutePrefix{
	Prefix: "/sales",
	SubRoutes: []customHTTP.Route{
		{
			Name:        "RegisterSale",
			Method:      http.MethodPost,
			Pattern:     "",
			HandlerFunc: CreateSale,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
		{
			Name:        "GetBuyers",
			Method:      http.MethodGet,
			Pattern:     "",
			HandlerFunc: GetSales,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
	},
}
