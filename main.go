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
	r.HandleFunc("/api/bookmarks/", DeleteUrlsHandler).Methods("DELETE")

	log.Print("Starting bookmark-service on " + port)

	log.Fatal(http.ListenAndServe(port, r))
}

func DeleteUrlsHandler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	urlsToDelete := queryValues["urls"]

	if len(urlsToDelete) == 0 {
		w.WriteHeader(400)
		return
	}
	log.Print(urlsToDelete)

	w.WriteHeader(200)
}
