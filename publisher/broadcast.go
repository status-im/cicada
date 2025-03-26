package publisher

import (
	"crypto/sha256"
	"time"

	"github.com/status-im/cicadian/feeds"
	broadcastpb "github.com/status-im/cicadian/proto"
)

// ToProto converts a FeedItem into a WakuFeedBroadcast protobuf message
func ToProto(item feeds.FeedItem) (*broadcastpb.WakuFeedBroadcast, error) {
	msg := &broadcastpb.WakuFeedBroadcast{
		Id:        item.ID,
		Title:     item.Title,
		Timestamp: item.Timestamp.Unix(),
		Link:      item.Link,
		ImageData: item.ImageData,
	}

	// Construct hash input using all core fields
	hashInput := item.ID + item.Title + item.Link + item.Timestamp.Format(time.RFC3339)
	hash := sha256.New()
	hash.Write([]byte(hashInput))
	hash.Write(item.ImageData) // Include image bytes
	msg.MessageHash = hash.Sum(nil)

	return msg, nil
}
