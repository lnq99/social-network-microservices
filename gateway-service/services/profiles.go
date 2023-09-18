package services

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewProfilesService() Service {
	target, _ := url.Parse(ProfilesServiceAddr)
	proxy := httputil.NewSingleHostReverseProxy(target)

	proxyHandler := func(p *httputil.ReverseProxy) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			p.ServeHTTP(w, r)
		}
	}(proxy)

	service := Service{
		Info: ServiceInfo{
			Name: "Profiles",
			Addr: ProfilesServiceAddr,
			Path: "",
		},
		Endpoints: []Endpoint{
			{"GET", "profile/{id}", proxyHandler},
			{"GET", "profile/short/{id}", proxyHandler},
			{"PATCH", "profile/intro", proxyHandler},

			{"GET", "rel/friends/{id}", proxyHandler},
			{"GET", "rel/mutual-friends/{id}", proxyHandler},
			{"GET", "rel/mutual-type/{id}", proxyHandler},
			{"PUT", "rel/{id}/{type}", proxyHandler},
		},
	}
	return service
}
