package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/benjih/bookmark-service/controllers"
	"github.com/benjih/bookmark-service/model"
)

func DeleteUrlsHandler(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	deleteRequest := &model.DeleteRequest{}
	if err = json.Unmarshal(requestBody, deleteRequest); err != nil {
		w.WriteHeader(400)
		return
	}
	if len(deleteRequest.URLs) > 0 {
		controllers.DeleteBookmarkByUrls(deleteRequest.URLs)
	}

	w.WriteHeader(200)
}
