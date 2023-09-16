package services

import "errors"

var (
	suStr                         = " service unavailable"
	ErrProfilesServiceUnavailable = errors.New("profiles" + suStr)
)

type ErrorResponse struct {
	Msg string `json:"message,omitempty"`
}

func ServiceUnavailable(service string) error {
	return errors.New(service + suStr)
}

func ToErrResponse(err error) ErrorResponse {
	return ErrorResponse{Msg: err.Error()}
}
