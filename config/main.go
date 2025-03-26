package config

type FeedConfig struct {
	RSS      []RSSFeedConfig      `yaml:"rss"`
	Twitter  TwitterConfig        `yaml:"twitter"`
	Ethereum []EthereumFeedConfig `yaml:"ethereum"`
}
