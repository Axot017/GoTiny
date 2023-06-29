package internal

import (
	"flag"
)

type Config struct {
	baseUrl        string
	logJson        bool
	linksTableName string
}

func NewConfig() *Config {
	cfg := Config{}
	flag.StringVar(&cfg.baseUrl, "base-url", "http://localhost:8080", "Base URL for shor links")
	flag.BoolVar(&cfg.logJson, "log-json", false, "User JSON format for logging")
	flag.StringVar(
		&cfg.linksTableName,
		"links-dynamodb-table",
		"go-tiny-dev",
		"DynamoDB table name for links",
	)
	flag.Parse()

	return &cfg
}

func (c *Config) BaseUrl() string {
	return c.baseUrl
}

func (c *Config) LogJson() bool {
	return c.logJson
}

func (c *Config) LinksTableName() string {
	return c.linksTableName
}
