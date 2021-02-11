package redeem

import (
	"net/http"

	customHTTP "github.com/openvino/openvino-api/src/http"
)

// Routes - Client related routes
var Routes = customHTTP.RoutePrefix{
	Prefix: "/redeem",
	SubRoutes: []customHTTP.Route{
		{
			Name:        "CreateRedeemInfo",
			Method:      http.MethodPost,
			Pattern:     "",
			HandlerFunc: CreateReedemInfo,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
		{
			Name:        "GetRedeemInfo",
			Method:      http.MethodGet,
			Pattern:     "",
			HandlerFunc: GetRedeemInfo,
			Scopes:      []customHTTP.Scope{customHTTP.WorkerScope},
		},
		{
			Name:        "GetShippingCosts",
			Method:      http.MethodGet,
			Pattern:     "/shipping",
			HandlerFunc: GetShippingCosts,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
	},
}
