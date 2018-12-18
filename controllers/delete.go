package controllers

import (
	"fmt"

	"github.com/benjih/bookmark-service/dao"
	"github.com/benjih/bookmark-service/model"
)

func DeleteBookmarkByUrls(urls []string) error {
	fmt.Println(urls)

	bookmarks, err := dao.RetrieveBookmarks()
	if err != nil {
		return err
	}

	deleteBookmarkFromTreeByURLs(bookmarks, urls)

	fmt.Println(bookmarks)

	dao.StoreBookmarks(bookmarks)
	return nil
}

func deleteBookmarkFromTreeByURLs(bookmarks []*model.Bookmark, urls []string) {
	bookmarksToDelete := []int{}
	for i, bookmark := range bookmarks {
		for _, url := range urls {
			if url == bookmark.URL {
				bookmarksToDelete = append(bookmarksToDelete, i)
				continue
			}

			if bookmark.Entries != nil {
				deleteBookmarkFromTreeByURLs(bookmark.Entries, urls)
			}
		}
	}

	for _, positionToDelete := range bookmarksToDelete {
		bookmarks = append(bookmarks[:positionToDelete], bookmarks[positionToDelete+1:]...)
	}
}
