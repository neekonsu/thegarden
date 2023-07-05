package main

import (
	"fmt"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func load(title string) (*Page, error) {
	
	var p *Page = &Page{title, nil}

	return p, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
