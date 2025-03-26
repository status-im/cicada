package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/status-im/cicadian/config"
	"github.com/status-im/cicadian/feeds"
	"github.com/status-im/cicadian/loader"

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
	err = wakuNode.Start(context.Background())
	if err != nil {
		log.Fatalf("Failed to start Waku node: %v", err)
	}
	defer wakuNode.Stop()

	// Set up feeds
	cfg, err := config.Read("feed_config.example.yaml")
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	fs, err := loader.Load(cfg)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
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
