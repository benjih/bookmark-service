package controllers

import (
	"github.com/benjih/bookmark-service/dao"
	"github.com/benjih/bookmark-service/model"
)

func DeleteBookmarkByUrls(urls []string) error {
	bookmarks, err := dao.RetrieveBookmarks()
	if err != nil {
		return err
	}

	bookmarks = deleteBookmarkFromTreeByURLs(bookmarks, urls)

	return dao.StoreBookmarks(bookmarks)
}

func deleteBookmarkFromTreeByURLs(bookmarks []*model.Bookmark, urls []string) []*model.Bookmark {
	bookmarksToDelete := []int{}
	for i, bookmark := range bookmarks {
		for _, url := range urls {
			if url == bookmark.URL {
				// TODO: This will delete a parent and all its entries

				// Define what behaviour should happen in this scenario
				bookmarksToDelete = append(bookmarksToDelete, i)
				continue
			}

			if bookmark.Entries != nil {
				bookmark.Entries = deleteBookmarkFromTreeByURLs(bookmark.Entries, urls)
			}
		}
	}

	var positionsDeleted = 0
	for _, positionToDelete := range bookmarksToDelete {
		adjustedPositionToDelete := positionToDelete - positionsDeleted
		bookmarks = append(bookmarks[:adjustedPositionToDelete], bookmarks[adjustedPositionToDelete+1:]...)
		positionsDeleted = positionsDeleted + 1
	}

	return bookmarks
}
