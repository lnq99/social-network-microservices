package services

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

//type SearchService interface {
//	GetSearch(id, limit, offset int) (feed []int64, err error)
//}

func NewSearchService() Service {
	target, _ := url.Parse(ProfilesServiceAddr)
	proxy := httputil.NewSingleHostReverseProxy(target)

	profilesProxyHandler := func(p *httputil.ReverseProxy) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			p.ServeHTTP(w, r)
		}
	}(proxy)

	service := Service{
		Info: ServiceInfo{
			Name: "Search",
			Addr: "",
			Path: "search",
		},
		Endpoints: []Endpoint{
			{"GET", "", profilesProxyHandler},
		},
	}
	return service
}

//func searchName(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode([]int{
//		1, 2, 3,
//	})
//}
