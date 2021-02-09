package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang/gddo/httputil/header"
	customHTTP "github.com/openvino/openvino-api/src/http"
)

// ContentType - Header name for content type
const ContentType = "Content-Type"

// ContentValue - Expected content type value
const ContentValue = "application/json"

// ContentTypeMiddleware - Checks if Content-Type header is valid
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(ContentType) != "" {
			value, _ := header.ParseValueAndParams(r.Header, ContentType)
			if value != ContentValue {
				customHTTP.NewErrorResponse(w, http.StatusUnsupportedMediaType, fmt.Sprintf("%s header must be %s", ContentType, ContentValue))
				return
			}
		}

		r.Body = http.MaxBytesReader(w, r.Body, 1048576)

		next.ServeHTTP(w, r)
	})
}
