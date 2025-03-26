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
	}

	// Compute a simple hash of key fields
	hash := sha256.Sum256([]byte(item.ID + item.Title + item.Timestamp.Format(time.RFC3339)))
	msg.MessageHash = hash[:]

	return msg, nil
}
