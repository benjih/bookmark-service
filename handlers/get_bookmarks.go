package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/benjih/bookmark-service/controllers"
)

func GetBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	bookmarkView, err := controllers.RetrieveBookmarks()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	jsonBytes, err := json.Marshal(bookmarkView)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Write(jsonBytes)
	w.WriteHeader(200)
}
