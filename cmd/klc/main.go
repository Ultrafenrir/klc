package main

import (
	"github.com/Ultrafenrir/klc/pkg/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/http-check", handlers.Hello)
	http.HandleFunc("/list-headers", handlers.Headers)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
