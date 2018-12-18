package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port = ":3000"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/bookmarks/__urls", DeleteUrlsHandler).Methods("DELETE")

	log.Print("Starting bookmark-service on " + port)

	log.Fatal(http.ListenAndServe(port, r))
}

func DeleteUrlsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
