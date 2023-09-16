package services

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httputil"
)

var (
	//ProfilesServiceAddr = ""
	Client = &http.Client{}
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
	Proxy     *httputil.ReverseProxy
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

//func chiProcessResponse[T any](c *chi.Ctx, status int, body io.ReadCloser, err error) error {
//	var r T
//
//	if err == nil {
//		err = json.NewDecoder(body).Decode(&r)
//		body.Close()
//	}
//
//	c.Status(status)
//	if err != nil {
//		return c.JSON(ToErrResponse(err))
//	}
//	return c.JSON(r)
//}
