package services

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewPostsService() Service {
	target, _ := url.Parse(PostsServiceAddr)
	proxy := httputil.NewSingleHostReverseProxy(target)

	proxyHandler := func(p *httputil.ReverseProxy) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			p.ServeHTTP(w, r)
		}
	}(proxy)

	service := Service{
		Info: ServiceInfo{
			Name: "Posts",
			Addr: PostsServiceAddr,
			Path: "",
		},
		Endpoints: []Endpoint{
			{"GET", "post/{id}", proxyHandler},
			{"GET", "post/u/{id}", proxyHandler},
			{"POST", "post", proxyHandler},
			{"DELETE", "post/{id}", proxyHandler},

			{"GET", "react/{post_id}", proxyHandler},
			{"GET", "react/u/{u_id}", proxyHandler},
			{"PUT", "react/{post_id}/{type}", proxyHandler},

			{"GET", "cmt/{id}", proxyHandler},
			{"POST", "cmt", proxyHandler},

			{"GET", "photo/{id}", proxyHandler},
			{"GET", "photo/u/{id}", proxyHandler},
		},
	}
	return service
}
