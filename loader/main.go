package loader

import (
	"github.com/status-im/cicadian/config"
	"github.com/status-im/cicadian/feeds"
)

func Load(cfg config.FeedConfig) ([]feeds.Feed, error) {
	var all []feeds.Feed
	all = append(all, LoadEthereumFeeds(cfg.Ethereum)...)
	all = append(all, LoadRSSFeeds(cfg.RSS)...)
	all = append(all, LoadTwitterFeeds(cfg.Twitter)...)

	// TODO: Add Reddit feed (JSON parsing)
	// TODO: Add Snapshot DAO proposal feed
	// TODO: Add Farcaster feed
	// TODO: Add Bluesky feed
	// TODO: Add Lens Protocol feed
	return all, nil
}
