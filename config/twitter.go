package config

type TwitterConfig struct {
	Credentials TwitterCredentials    `yaml:"credentials"`
	Targets     []TwitterTargetConfig `yaml:"targets"`
}

type TwitterCredentials struct {
	ConsumerKeyEnv    string `yaml:"consumer_key_env"`
	ConsumerSecretEnv string `yaml:"consumer_secret_env"`
	AccessTokenEnv    string `yaml:"access_token_env"`
	AccessSecretEnv   string `yaml:"access_secret_env"`
}

type TwitterTargetType string

const (
	TwitterProfile TwitterTargetType = "profile"
	TwitterSearch  TwitterTargetType = "search"
	TwitterHashtag TwitterTargetType = "hashtag"
	TwitterList    TwitterTargetType = "list"
)

type TwitterTargetConfig struct {
	Name string            `yaml:"name"`
	Type TwitterTargetType `yaml:"type"`
}
