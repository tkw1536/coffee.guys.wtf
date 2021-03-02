package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static
var staticFS embed.FS

// the root filesysem
var rootFS fs.FS

func init() {
	var err error
	rootFS, err = fs.Sub(staticFS, "static")
	if err != nil {
		panic(err)
	}
}

var notFound = []byte("Not Found")
var internalServerError = []byte("Internal Server Error")

func main() {

	server := http.FileServer(http.FS(rootFS))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "", "/", "/index.html":
			bytes, err := fs.ReadFile(rootFS, "index.html")
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
