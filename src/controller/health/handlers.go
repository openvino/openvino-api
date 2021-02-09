package health

import (
	"net/http"

	customHTTP "github.com/openvino/openvino-api/src/http"
)

type healthResponse struct {
	Status string `json:"status"`
}

// Handler - Write new value
func Handler(w http.ResponseWriter, r *http.Request) {
	var response healthResponse
	response.Status = "UP"

	customHTTP.ResponseJSON(w, response)
}
