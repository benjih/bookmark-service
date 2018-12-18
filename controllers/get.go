package controllers

import (
	"github.com/benjih/bookmark-service/dao"
	"github.com/benjih/bookmark-service/model"
)

func RetrieveBookmarks() (*model.BookmarkView, error) {
	bookmarks, err := dao.RetrieveBookmarks()
	if err != nil {
		return nil, err
	}

	return &model.BookmarkView{
		Data: bookmarks,
	}, nil
}
