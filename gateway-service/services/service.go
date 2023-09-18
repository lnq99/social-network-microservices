package services

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

var (
	ProfilesServiceAddr = ""
	PostsServiceAddr    = ""
	StatsServiceAddr    = ""
	Client              = &http.Client{}
	//AuthHeader          = "Authorization"
	//UsernameHeader      = "Name"
)

type Endpoint struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type ServiceInfo struct {
	Name string
	Addr string
	Path string
}

type Service struct {
	Info      ServiceInfo
	Endpoints []Endpoint
}

type ChiRouter struct {
	Router chi.Router
}

func (s ChiRouter) RegisterRoute(e *Endpoint, prefix string) {
	s.Router.Method(e.Method, prefix+e.Path, e.Handler)
}

func (s ChiRouter) RegisterService(service Service) {
	prefix := "/" + service.Info.Path
	for _, e := range service.Endpoints {
		s.RegisterRoute(&e, prefix)
	}
}

func processResponse[T any](w http.ResponseWriter, status int, body io.ReadCloser, err error) error {
	var r T

	if err == nil {
		err = json.NewDecoder(body).Decode(&r)
		body.Close()
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		err = json.NewEncoder(w).Encode(ToErrResponse(err))
		return err
	}
	err = json.NewEncoder(w).Encode(r)
	return err
}
