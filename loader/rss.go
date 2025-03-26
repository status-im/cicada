package loader

import (
	"github.com/status-im/cicadian/config"
	"github.com/status-im/cicadian/feeds"
)

func LoadRSSFeeds(configs []config.RSSFeedConfig) []feeds.Feed {
	var result []feeds.Feed
	for _, cfg := range configs {
		result = append(result, feeds.NewRSSFeed(cfg.URL))
	}
	return result
}
