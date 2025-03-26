package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type FeedConfig struct {
	Ethereum []EthereumFeedConfig `yaml:"ethereum"`
	RSS      []RSSFeedConfig      `yaml:"rss"`
	Snapshot []SnapshotFeedConfig `yaml:"snapshot"`
	Twitter  TwitterConfig        `yaml:"twitter"`
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
