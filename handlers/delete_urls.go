package handlers

import (
	"log"
	"net/http"
)

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
