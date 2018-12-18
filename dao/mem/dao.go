package mem

import "github.com/benjih/bookmark-service/model"

type InMemoryDAO struct {
	bookmarks []*model.Bookmark
}

func NewInMemoryDAO() *InMemoryDAO {
	return &InMemoryDAO{
		bookmarks: []*model.Bookmark{},
	}
}

func (d *InMemoryDAO) StoreBookmarks(bookmarks []*model.Bookmark) error {
	d.bookmarks = bookmarks
	return nil
}
func (d *InMemoryDAO) RetrieveBookmarks() ([]*model.Bookmark, error) {
	return d.bookmarks, nil
}
