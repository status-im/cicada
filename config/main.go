package config

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type FeedConfig struct {
	RSS      []RSSFeedConfig      `yaml:"rss"`
	Twitter  TwitterConfig        `yaml:"twitter"`
	Ethereum []EthereumFeedConfig `yaml:"ethereum"`
}

func Read(path string) (FeedConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return FeedConfig{}, err
	}

	var cfg FeedConfig
	err = yaml.Unmarshal(data, &cfg)
	return cfg, err
}
