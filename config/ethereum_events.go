package config

type EthereumFeedConfig struct {
	RPCUrl     string `yaml:"rpc_url"`
	Contract   string `yaml:"contract"`
	Event      string `yaml:"event"`
	StartBlock uint64 `yaml:"start_block"`
}
