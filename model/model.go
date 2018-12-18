package model

type BookmarkView struct {
	Data []*Bookmark `json:"data"`
}

type Bookmark struct {
	Name    string      `json:"name"`
	URL     string      `json:"url"`
	Entries []*Bookmark `json:"entries,omitempty"`
}

type DeleteRequest struct {
	URLs []string
}
