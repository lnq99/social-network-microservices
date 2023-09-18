package services

import (
	"encoding/json"
	"net/http"
)

type SearchService interface {
	GetSearch(id, limit, offset int) (feed []int64, err error)
}

func NewSearchService() Service {
	service := Service{
		Info: ServiceInfo{
			Name: "Search",
			Addr: "",
			Path: "",
		},
		Endpoints: []Endpoint{
			{"GET", "search", getSearch},
		},
	}
	return service
}

func getSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]int{
		1, 2, 3,
	})
}
