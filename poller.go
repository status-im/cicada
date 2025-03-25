package main

import (
	"log"

	"github.com/status-im/cicadian/feeds"
)

func PollFeed(feed feeds.Feed, seen map[string]bool, publish func(feeds.FeedItem)) {
	items, err := feed.FetchItems()
	if err != nil {
		log.Printf("Error fetching from %s: %v", feed.Name(), err)
		return
	}

	for _, item := range items {
		if !seen[item.ID] {
			seen[item.ID] = true
			publish(item)
		}
	}
}
