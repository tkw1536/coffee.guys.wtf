package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var notFound = []byte("Not Found")
var internalServerError = []byte("Internal Server Error")

func main() {

	server := http.FileServer(http.Dir("./static"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "", "/", "/index.html":
			bytes, err := ioutil.ReadFile("./static/index.html")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(internalServerError)
				return
			}
			w.WriteHeader(418)
			w.Write(bytes)
		default:
			server.ServeHTTP(w, r)
		}
	})
	fmt.Println("coffee.guys.wtf listening on 0.0.0.0:8080")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal(err)
	}
}
