package feeds

import "time"

type FeedItem struct {
	ID        string
	Title     string
	Link      string
	Timestamp time.Time
	//TODO: ImageData []byte // Add this later
}

type Feed interface {
	Name() string
	FetchItems() ([]FeedItem, error)
}
