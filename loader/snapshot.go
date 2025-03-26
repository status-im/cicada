package loader

import (
	"github.com/status-im/cicadian/config"
	"github.com/status-im/cicadian/feeds"
)

func LoadSnapshotFeeds(configs []config.SnapshotFeedConfig) []feeds.Feed {
	var result []feeds.Feed
	for _, cfg := range configs {
		result = append(result, feeds.NewSnapshotFeed(cfg.SpaceID))
	}
	return result
}
