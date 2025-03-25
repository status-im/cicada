package cicadian

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
	waku "github.com/waku-org/go-waku/waku/v2/node"
	"github.com/waku-org/go-waku/waku/v2/protocol/pb"
)

const (
	rssFeedURL       = "https://blog.waku.org/rss/"
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

	// Ensure the node stops when the application exits
	defer wakuNode.Stop()

	parser := gofeed.NewParser()
	seen := make(map[string]bool)

	for {
		feed, err := parser.ParseURL(rssFeedURL)
		if err != nil {
			log.Println("Error fetching RSS feed:", err)
			continue
		}

		for _, item := range feed.Items {
			if !seen[item.GUID] {
				seen[item.GUID] = true
				publishToWaku(ctx, wakuNode, item)
			}
		}

		time.Sleep(pollInterval)
	}
}

func publishToWaku(ctx context.Context, node *waku.WakuNode, item *gofeed.Item) {
	payload := fmt.Sprintf("Title: %s\nLink: %s\nPublished: %s", item.Title, item.Link, item.Published)

	ts := time.Now().UnixNano()
	msg := &pb.WakuMessage{
		Payload:      []byte(payload),
		ContentTopic: wakuContentTopic,
		Version:      new(uint32),
		Timestamp:    &ts,
	}

	_, err := node.Relay().Publish(ctx, msg)
	if err != nil {
		log.Println("Failed to publish to Waku:", err)
	} else {
		log.Println("Published to Waku:", item.Title)
	}
}
