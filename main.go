package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
)

type Page struct {
	Title 	string
	body	[]byte
}

func (p *Page) savePage() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}