package main

import (
	"log"
	"net/http"

	"github.com/benjih/bookmark-service/handlers"
	"github.com/gorilla/mux"
)

const (
	port = ":3000"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/bookmarks/", handlers.DeleteUrlsHandler).Methods("DELETE")

	log.Print("Starting bookmark-service on " + port)

	log.Fatal(http.ListenAndServe(port, r))
}
