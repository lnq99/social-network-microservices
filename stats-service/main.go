package main

import (
	"log"
	"net/http"

	"app/config"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	log.Fatal(http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.Port, nil))
}
