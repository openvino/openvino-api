package sensor

import (
	"net/http"

	customHTTP "github.com/openvino/openvino-api/src/http"
)

// Routes - Client related routes
var Routes = customHTTP.RoutePrefix{
	Prefix: "/sensor_data",
	SubRoutes: []customHTTP.Route{
		{
			Name:        "SaveSensorData",
			Method:      http.MethodPost,
			Pattern:     "",
			HandlerFunc: SaveSensorRecords,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
		{
			Name:        "GetSensorData",
			Method:      http.MethodGet,
			Pattern:     "",
			HandlerFunc: GetSensorRecords,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
		{
			Name:        "GetHashes",
			Method:      http.MethodGet,
			Pattern:     "/hashes",
			HandlerFunc: GetSensorHashes,
			Scopes:      []customHTTP.Scope{customHTTP.GuestScope},
		},
	},
}
