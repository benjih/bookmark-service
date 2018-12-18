package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/benjih/bookmark-service/controllers"
	"github.com/benjih/bookmark-service/model"
)

func StoreBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	bookmarkView := &model.BookmarkView{}
	if err = json.Unmarshal(requestBody, bookmarkView); err != nil {
		w.WriteHeader(400)
		return
	}

	if err = controllers.StoreBookmarks(bookmarkView); err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}
