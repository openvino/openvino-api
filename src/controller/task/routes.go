package task

import (
	"net/http"

	customHTTP "github.com/openvino/openvino-api/src/http"
)

// Routes - Client related routes
var Routes = customHTTP.RoutePrefix{
	Prefix: "/tasks",
	SubRoutes: []customHTTP.Route{
		{
			Name:        "CreateTask",
			Method:      http.MethodPost,
			Pattern:     "",
			HandlerFunc: CreateTask,
			Scopes:      []customHTTP.Scope{customHTTP.WorkerScope},
		},
		{
			Name:        "GetTasks",
			Method:      http.MethodGet,
			Pattern:     "",
			HandlerFunc: GetTasks,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
	},
}
