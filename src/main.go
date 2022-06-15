package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/createIntent", createIntent)
	mux.HandleFunc("/replyIntent", replyIntent)
	// mux.HandleFunc("/about", about)
	port := ":9000"
	log.Println("Listening on port ", port)
	http.ListenAndServe(port, mux)
}
