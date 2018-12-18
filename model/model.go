package model

type BookmarkView struct {
	Data []*Bookmark
}

type Bookmark struct {
	Name    string
	URL     string
	Entries []*Bookmark
}
