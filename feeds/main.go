package feeds

import "time"

type FeedItem struct {
	ID        string
	Title     string
	Link      string
	Timestamp time.Time
}

type Feed interface {
	Name() string
	FetchItems() ([]FeedItem, error)
}
