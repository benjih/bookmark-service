package dao

import "github.com/benjih/bookmark-service/model"

type DAO interface {
	StoreBookmarks([]*model.Bookmark) error
	RetrieveBookmarks() ([]*model.Bookmark, error)
}
