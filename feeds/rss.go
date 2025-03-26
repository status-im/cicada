package feeds

import (
	"github.com/mmcdole/gofeed"
)

type RSSFeed struct {
	url    string
	parser *gofeed.Parser
	seen   map[string]bool
}

func NewRSSFeed(url string) Feed {
	return &RSSFeed{
		url:    url,
		parser: gofeed.NewParser(),
		seen:   make(map[string]bool),
	}
}

func (r *RSSFeed) Name() string {
	return "rss:" + r.url
}

func (r *RSSFeed) FetchItems() ([]FeedItem, error) {
	feed, err := r.parser.ParseURL(r.url)
	if err != nil {
		return nil, err
	}

	var items []FeedItem
	for _, entry := range feed.Items {
		if !r.seen[entry.GUID] {
			r.seen[entry.GUID] = true
			items = append(items, FeedItem{
				ID:        entry.GUID,
				Title:     entry.Title,
				Link:      entry.Link,
				Timestamp: *entry.PublishedParsed,
			})
		}
	}
	return items, nil
}
