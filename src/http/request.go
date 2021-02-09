package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/thedevsaddam/govalidator"
)

type malformedRequest struct {
	msg string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}

// DecodeJSONBody - Decodes a JSON body into an interface. Sends an HTTP response if an error occurs
func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}, rules map[string][]string) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			NewErrorResponse(w, http.StatusBadRequest, msg)
			return &malformedRequest{msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			NewErrorResponse(w, http.StatusBadRequest, msg)
			return &malformedRequest{msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			NewErrorResponse(w, http.StatusBadRequest, msg)
			return &malformedRequest{msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			NewErrorResponse(w, http.StatusBadRequest, msg)
			return &malformedRequest{msg}

		case errors.Is(err, io.EOF):
			break

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			NewErrorResponse(w, http.StatusBadRequest, msg)
			return &malformedRequest{msg}

		default:
			NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		NewErrorResponse(w, http.StatusInternalServerError, msg)
		return &malformedRequest{msg}
	}

	validateBody(r, rules, &dst)

	return nil
}

// ValidateBody - Validate body based on rules
func validateBody(r *http.Request, rules map[string][]string, data interface{}) url.Values {
	opts := govalidator.Options{
		Request: r,
		Data:    &data,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	return e
}
