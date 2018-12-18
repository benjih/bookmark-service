package controllers

import (
	"github.com/benjih/bookmark-service/dao"
	"github.com/benjih/bookmark-service/model"
)

func StoreBookmarks(bookmarkView *model.BookmarkView) error {
	return dao.StoreBookmarks(bookmarkView.Data)
}
