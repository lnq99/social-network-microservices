package services

import (
	"context"
	"io"
	"net/http"
	"time"

	breaker "github.com/sony/gobreaker"
)

var (
	CbSetting = breaker.Settings{
		Name:        "Default",
		MaxRequests: 1,
		Interval:    10 * time.Second,
		Timeout:     2 * time.Second,
		ReadyToTrip: func(counts breaker.Counts) bool {
			return counts.ConsecutiveFailures > 4
		},
		OnStateChange: nil,
		IsSuccessful:  nil,
	}

	profilesCb = breaker.NewCircuitBreaker(CbSetting)
	postsCb    = breaker.NewCircuitBreaker(CbSetting)
)

type ServiceResponse struct {
	status int
	body   io.ReadCloser
}

func CallServiceWithCircuitBreaker(
	cb *breaker.CircuitBreaker, method, url string,
	header http.Header, body io.Reader) (ServiceResponse, error) {

	res, err := cb.Execute(func() (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		req, _ := http.NewRequestWithContext(ctx, method, url, body)
		req.Header = header
		res, err := Client.Do(req)
		if err != nil {
			return ServiceResponse{status: http.StatusServiceUnavailable},
				ErrProfilesServiceUnavailable
		}
		return ServiceResponse{
			status: res.StatusCode,
			body:   res.Body,
		}, nil
	})

	if err == breaker.ErrOpenState {
		return ServiceResponse{}, err
	}

	return res.(ServiceResponse), err
}
