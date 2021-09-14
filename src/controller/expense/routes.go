package expense

import (
	"net/http"

	customHTTP "github.com/openvino/openvino-api/src/http"
)

// Routes - Client related routes
var Routes = customHTTP.RoutePrefix{
	Prefix: "/expenses",
	SubRoutes: []customHTTP.Route{
		{
			Name:        "CreateExpense",
			Method:      http.MethodPost,
			Pattern:     "",
			HandlerFunc: CreateExpense,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
		{
			Name:        "GetExpenses",
			Method:      http.MethodGet,
			Pattern:     "",
			HandlerFunc: GetExpenses,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
	},
}
