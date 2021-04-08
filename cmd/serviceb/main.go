package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("SERVICE_B_PORT")
	http.HandleFunc("/hello", HelloHandler)
	log.Printf("listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Service B\n"))
}
