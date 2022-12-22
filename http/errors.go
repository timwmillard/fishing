package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/timwmillard/fishing"
)

type ErrorType string

const (
	ServerError     ErrorType = "server_error"
	BadRequestError ErrorType = "bad_request"
	NotFoundError   ErrorType = "not_found"
)

type Error struct {
	Type        ErrorType `json:"type"`
	Description string    `json:"description,omitempty"`
	httpCode    int       `json:"-"`
}

func (e Error) Error() string {
	return e.Description
}

func handleJSONDecodeError(wr http.ResponseWriter, err error, action string) {
	log.Printf("%s json decode error: %v", action, err)
	handleError(wr, fromJSONDecodeError(err), action)
}

func handleError(wr http.ResponseWriter, err error, action string) {
	servErr := fromError(err)
	log.Printf("%s error: %v", action, err)
	wr.WriteHeader(servErr.httpCode)
	errJson := json.NewEncoder(wr).Encode(servErr)
	if errJson != nil {
		log.Printf("%s: encoding error: %v", action, errJson)
	}
}

func fromError(err error) Error {
	switch {
	case errors.Is(err, fishing.ErrCompetitorNotFound):
		return Error{
			Type:        NotFoundError,
			Description: "Competitor not found",
			httpCode:    http.StatusNotFound,
		}
	default:
		return Error{
			Type:        ServerError,
			Description: "Server",
			httpCode:    http.StatusInternalServerError,
		}
	}
}

func fromJSONDecodeError(err error) Error {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	switch {
	case errors.As(err, &syntaxError):
		return Error{
			Type:        BadRequestError,
			Description: fmt.Sprintf("body contains badly-formed JSON (at character %d)", syntaxError.Offset),
			httpCode:    http.StatusBadRequest,
		}
	case errors.Is(err, io.ErrUnexpectedEOF):
		return Error{
			Type:        BadRequestError,
			Description: "body contains badly-formed JSON",
			httpCode:    http.StatusBadRequest,
		}
	case errors.As(err, &unmarshalTypeError):
		var desc string
		if unmarshalTypeError.Field != "" {
			desc = fmt.Sprintf("body contains incorrect JSON type for field '%s'", unmarshalTypeError.Field)
		} else {
			desc = fmt.Sprintf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		}
		return Error{
			Type:        BadRequestError,
			Description: desc,
			httpCode:    http.StatusBadRequest,
		}
	case errors.Is(err, io.EOF):
		return Error{
			Type:        BadRequestError,
			Description: "body must not be empty",
			httpCode:    http.StatusBadRequest,
		}
	default:
		return Error{
			Type:        BadRequestError,
			Description: "body decoding error",
			httpCode:    http.StatusBadRequest,
		}
	}
}
