package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	targetURL, port := os.Getenv("SERVICE_A_URL"), os.Getenv("API_GATEWAY_PORT")
	target, err := url.Parse(targetURL)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/hello", httputil.NewSingleHostReverseProxy(target))
	log.Printf("listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
