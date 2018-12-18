package dao

import (
	"github.com/benjih/bookmark-service/dao/mem"
	"github.com/benjih/bookmark-service/model"
)

type DAO interface {
	StoreBookmarks(bookmarks []*model.Bookmark) error
	RetrieveBookmarks() ([]*model.Bookmark, error)
}

var GlobalDAO = mem.NewInMemoryDAO()

func StoreBookmarks(bookmarks []*model.Bookmark) error {
	return GlobalDAO.StoreBookmarks(bookmarks)
}
func RetrieveBookmarks() ([]*model.Bookmark, error) {
	return GlobalDAO.RetrieveBookmarks()
}
