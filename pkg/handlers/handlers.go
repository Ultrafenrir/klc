package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {

	_, err := fmt.Fprintf(w, "hello\n")
	if err != nil {
		log.Fatal(err)
	}
}

func Headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for h := range headers {
			_, err := fmt.Fprintf(w, "%v: %v\n", name, h)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
