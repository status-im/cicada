package loader

import (
	"log"
	"os"

	"github.com/status-im/cicadian/config"
	"github.com/status-im/cicadian/feeds"
)

func LoadTwitterFeeds(cfg config.TwitterConfig) []feeds.Feed {
	var result []feeds.Feed

	client := feeds.NewTwitterClient(
		os.Getenv(cfg.Credentials.ConsumerKeyEnv),
		os.Getenv(cfg.Credentials.ConsumerSecretEnv),
		os.Getenv(cfg.Credentials.AccessTokenEnv),
		os.Getenv(cfg.Credentials.AccessSecretEnv),
	)

	for _, target := range cfg.Targets {
		switch target.Type {
		case config.TwitterProfile:
			result = append(result, feeds.NewTwitterProfileFeed(client.Client(), target.Name))
		case config.TwitterSearch:
			result = append(result, feeds.NewTwitterSearchFeed(client.Client(), target.Name))
		case config.TwitterHashtag:
			result = append(result, feeds.NewTwitterSearchFeed(client.Client(), "#"+target.Name))
		case config.TwitterList:
			log.Printf("twitter:list type is defined but not yet supported (target: %s)", target.Name)

		default:
			log.Printf("unknown twitter target type: %s", target.Type)
		}
	}

	return result
}
