package feeds

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/mmcdole/gofeed"
	"strings"
	"time"
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
		if r.seen[entry.GUID] {
			continue
		}
		r.seen[entry.GUID] = true

		i, err := extractImageData(entry)
		if err != nil {
			log.Error("failed to fetch image", "link", entry.Link, "error", err)
		}

		ts := time.Now()
		if entry.PublishedParsed != nil {
			ts = *entry.PublishedParsed
		}

		items = append(items, FeedItem{
			ID:        entry.GUID,
			Title:     entry.Title,
			Link:      entry.Link,
			Timestamp: ts,
			ImageData: i,
		})
	}
	return items, nil
}

func extractImageData(entry *gofeed.Item) ([]byte, error) {
	var imageURL string

	if entry.Image != nil && entry.Image.URL != "" {
		imageURL = entry.Image.URL
	} else {
		for _, enc := range entry.Enclosures {
			if strings.HasPrefix(enc.Type, "image/") {
				imageURL = enc.URL
				break
			}
		}
	}

	// Fallback: parse Markdown or HTML from content/description
	if imageURL == "" {
		imageURL = extractFirstImageURL(entry.Content)
	}
	if imageURL == "" {
		imageURL = extractFirstImageURL(entry.Description)
	}

	if imageURL == "" {
		return nil, nil
	}

	return fetchImageBytes(imageURL)
}
