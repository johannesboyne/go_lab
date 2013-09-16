package main

import (
	"fmt"
	"net/http"
)

type Page struct {
	Title string
	Body  string
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "Welt", Body: "Experimenting with Go for fun and profit."}
	fmt.Fprintf(w, "<html><head><title>%s</title></head><body>%s<br/>Hi there, I love %s!</body>", p.Title, p.Body, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
