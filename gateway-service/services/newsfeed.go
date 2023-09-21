package services

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"strconv"
	"strings"
)

type NewsFeedService interface {
	GetNewsFeed(id, limit, offset int) (feed []int64, err error)
}

func NewNewsFeedService() Service {
	//target, _ := url.Parse(PostsServiceAddr)
	//proxy := httputil.NewSingleHostReverseProxy(target)
	//
	//postsProxyHandler := func(p *httputil.ReverseProxy) http.HandlerFunc {
	//	return func(w http.ResponseWriter, r *http.Request) {
	//		p.ServeHTTP(w, r)
	//	}
	//}(proxy)

	service := Service{
		Info: ServiceInfo{
			Name: "NewsFeed",
			Addr: "",
			Path: "feed",
		},
		Endpoints: []Endpoint{
			{"GET", "", getNewsFeed},
		},
	}
	return service
}

type friendsResponse struct {
	ID int `json:"id"`
}

func getNewsFeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode([]int{
	//	1, 2, 3,
	//})

	_, claims, _ := jwtauth.FromContext(r.Context())

	var reqUrl string

	reqUrl = fmt.Sprintf("%s/api/v1/rel/friends/%d",
		ProfilesServiceAddr, int(claims["ID"].(float64)))

	res, _ := CallServiceWithCircuitBreaker(
		profilesCb, "GET", reqUrl, r.Header, nil)

	var reqBody1 []friendsResponse
	_ = json.NewDecoder(res.body).Decode(&reqBody1)
	_ = r.Body.Close()

	ids_arr := make([]string, len(reqBody1))
	for i, v := range reqBody1 {
		ids_arr[i] = strconv.Itoa(v.ID)
	}

	lim := r.URL.Query().Get("lim")
	off := r.URL.Query().Get("off")

	reqUrl = fmt.Sprintf("%s/api/v1/feed?lim=%s&off=%s&ids=%s",
		PostsServiceAddr,
		lim,
		off,
		strings.Join(ids_arr, ","))

	res, _ = CallServiceWithCircuitBreaker(
		profilesCb, "GET", reqUrl, r.Header, nil)

	var reqBody2 interface{}
	_ = json.NewDecoder(res.body).Decode(&reqBody2)
	_ = r.Body.Close()

	w.WriteHeader(res.status)
	json.NewEncoder(w).Encode(reqBody2)
}
