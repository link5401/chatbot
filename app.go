package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./tpl/*.html",
	}

	rnd = renderer.New(opts)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	// mux.HandleFunc("/about", about)
	port := ":9000"
	log.Println("Listening on port ", port)
	http.ListenAndServe(port, mux)
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
