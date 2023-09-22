package services

import (
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

func NewStatsService() Service {
	//target, _ := url.Parse(StatsServiceAddr)
	//proxy := httputil.NewSingleHostReverseProxy(target)
	//
	//statsProxyHandler := func(p *httputil.ReverseProxy) http.HandlerFunc {
	//	return func(w http.ResponseWriter, r *http.Request) {
	//		p.ServeHTTP(w, r)
	//	}
	//}(proxy)

	service := Service{
		Info: ServiceInfo{
			Name: "Stats",
			Addr: "",
			Path: "",
		},
		Endpoints: []Endpoint{
			{"GET", "admin/log", getLogs},
		},
	}
	return service
}

func getLogs(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "loginHandler, %q", html.EscapeString(r.URL.Path))

	var err error
	var roleStr string
	//var account repository.Account

	_, claims, err := jwtauth.FromContext(r.Context())
	fmt.Printf("%+v\n", claims)

	if err == nil {
		fmt.Printf("%+v\n", claims)

		role, ok := claims["role"]
		if ok {
			roleStr = role.(string)
		}

	} else {
		fmt.Println(err)
	}

	if roleStr == "admin" {
		w.Write([]byte(fmt.Sprintf("protected area. hi %+v", claims)))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

	//processResponse[interface{}](w, req.status, req.body, err)
}
