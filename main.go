package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed static/index.html
var indexHTML []byte
var notFound = []byte("Not Found")

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "", "/", "/index.html":
			w.WriteHeader(http.StatusTeapot)
			w.Header().Add("Content-Type", "text/html")
			w.Write(indexHTML)
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Header().Add("Content-Type", "text/plain")
			w.Write(notFound)
		}
	})
	fmt.Println("coffee.guys.wtf listening on 0.0.0.0:8080")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal(err)
	}
}
