package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/status-im/cicadian/feeds"

	waku "github.com/waku-org/go-waku/waku/v2/node"
	"github.com/waku-org/go-waku/waku/v2/protocol/pb"
)

const (
	wakuContentTopic = "announcements"
	pollInterval     = 5 * time.Minute
)

func main() {
	ctx := context.Background()

	// Create a new Waku node with the Relay protocol enabled
	wakuNode, err := waku.New(waku.WithWakuRelay())
	if err != nil {
		log.Fatalf("Failed to create Waku node: %v", err)
	}

	// Start the Waku node
	if err := wakuNode.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start Waku node: %v", err)
	}
	defer wakuNode.Stop()

	// Set up feeds
	// TODO these are a bit hardcoded but we can make this more flexible later
	fs := []feeds.Feed{
		feeds.NewRSSFeed("https://blog.waku.org/rss/"),
		feeds.NewRSSFeed("https://github.com/waku-org/go-waku/releases.atom"),
		feeds.NewTwitterFeed(
			os.Getenv("TWITTER_CONSUMER_KEY"),
			os.Getenv("TWITTER_CONSUMER_SECRET"),
			os.Getenv("TWITTER_ACCESS_TOKEN"),
			os.Getenv("TWITTER_ACCESS_SECRET"),
			"Waku_org"),
		// TODO: Add Youtube feed (RSS)
		// TODO: Add Reddit feed (JSON parsing)
		// TODO: Add ethereum contract events feed
		// TODO: Add Snapshot DAO proposal feed
		// TODO: Add Farcaster feed
		// TODO: Add Bluesky feed
		// TODO: Add Lens Protocol feed
	}
	seen := make(map[string]bool)

	for {
		for _, feed := range fs {
			PollFeed(feed, seen, func(item feeds.FeedItem) {
				publishToWaku(ctx, wakuNode, item)
			})
		}
		time.Sleep(pollInterval)
	}
}

func publishToWaku(ctx context.Context, node *waku.WakuNode, item feeds.FeedItem) {
	payload := fmt.Sprintf("Title: %s\nLink: %s\nPublishedAt: %s", item.Title, item.Link, item.Timestamp)

	ts := time.Now().UnixNano()
	msg := &pb.WakuMessage{
		Payload:      []byte(payload),
		ContentTopic: wakuContentTopic,
		Version:      new(uint32),
		Timestamp:    &ts,
	}

	msgHash, err := node.Relay().Publish(ctx, msg)
	if err != nil {
		log.Println("Failed to publish to Waku:", err)
	} else {
		log.Println("Published to Waku:", item.Title, msgHash)
	}
}
