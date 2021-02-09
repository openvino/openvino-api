package util

import (
	customHTTP "github.com/openvino/openvino-api/src/http"
)

func ContainsScope(scope string, list []customHTTP.Scope) bool {
	for _, accepted := range list {
		if scope == accepted.String() {
			return true
		}
	}
	return false
}
