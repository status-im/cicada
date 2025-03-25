package feeds

import "github.com/mmcdole/gofeed"

type RSSFeed struct {
	URL    string
	parser *gofeed.Parser
}

func NewRSSFeed(url string) *RSSFeed {
	return &RSSFeed{
		URL:    url,
		parser: gofeed.NewParser(),
	}
}

func (r *RSSFeed) Name() string {
	return r.URL
}

func (r *RSSFeed) FetchItems() ([]FeedItem, error) {
	rawFeed, err := r.parser.ParseURL(r.URL)
	if err != nil {
		return nil, err
	}

	var items []FeedItem
	for _, i := range rawFeed.Items {
		items = append(items, FeedItem{
			ID:        i.GUID,
			Title:     i.Title,
			Link:      i.Link,
			Timestamp: *i.PublishedParsed,
		})
	}
	return items, nil
}
