package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("SERVICE_A_PORT")

	http.HandleFunc("/hello", HelloHandler)
	log.Printf("listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	b, err := CallServiceB(os.Getenv("SERVICE_B_URL") + "/hello")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error to call service b"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Service A and " + string(b)))
}

func CallServiceB(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	}

	return body, nil
}
