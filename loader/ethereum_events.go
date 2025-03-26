package loader

import (
	"log"

	"github.com/status-im/cicadian/config"
	"github.com/status-im/cicadian/feeds"
)

func LoadEthereumFeeds(configs []config.EthereumFeedConfig) []feeds.Feed {
	var result []feeds.Feed

	for _, cfg := range configs {
		feed, err := feeds.NewEthereumEventFeed(
			cfg.RPCUrl,
			cfg.Contract,
			cfg.Event,
			cfg.StartBlock,
		)
		if err != nil {
			log.Printf("Failed to init ethereum feed for contract %s: %v", cfg.Contract, err)
			continue
		}
		result = append(result, feed)
	}

	return result
}
