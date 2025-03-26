package main

import (
	"context"
	"log"
	"time"

	"github.com/status-im/cicadian/config"
	"github.com/status-im/cicadian/feeds"
	"github.com/status-im/cicadian/loader"
	"github.com/status-im/cicadian/publisher"

	waku "github.com/waku-org/go-waku/waku/v2/node"
	"github.com/waku-org/go-waku/waku/v2/protocol/pb"
	"google.golang.org/protobuf/proto"
)

const (
	wakuContentTopic = "cicadian_announcements"
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
	msg, err := publisher.ToProto(item)
	if err != nil {
		log.Printf("ToProto error: %v", err)
		return
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Printf("MarshalBroadcast error: %v", err)
		return
	}

	v := uint32(0)
	ts := time.Now().UnixNano()
	wakuMsg := &pb.WakuMessage{
		Payload:      data,
		ContentTopic: wakuContentTopic,
		Version:      &v,
		Timestamp:    &ts,
	}

	msgHash, err := node.Relay().Publish(ctx, wakuMsg)
	if err != nil {
		log.Println("Failed to publish to Waku:", err)
	} else {
		log.Println("Published to Waku:", item.Title, msgHash)
	}
}
