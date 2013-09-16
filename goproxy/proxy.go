package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	pH := http.HandlerFunc(proxyHandler)
	http.ListenAndServe(":8080", pH)
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// "business logic" log everything
	log.Println(r.URL)

	// new http client for each incoming request
	client := &http.Client{}

	// "proxy"
	req, _ := http.NewRequest("GET", "http://localhost:8081"+r.URL.String(), nil)
	resp, _ := client.Do(req)
	// what's defer?
	// => http://blog.golang.org/defer-panic-and-recover
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		fmt.Fprintf(w, bodyString)
	}
}
