package app

import "github.com/core-go/log/zap"

type Config struct {
	ElasticSearch ElasticSearchConfig `mapstructure:"elastic_search"`
	Log           log.Config          `mapstructure:"log"`
}

type ElasticSearchConfig struct {
	Url      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
