package internal

import (
	"os"
)

type Config struct {
	baseUrl        string
	logJson        bool
	linksTableName string
}

func NewConfig() *Config {
	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost:8080"
	}
	logJson := os.Getenv("LOG_JSON") == "true"
	linksTableName := os.Getenv("LINKS_DYNAMODB_TABLE")
	if linksTableName == "" {
		linksTableName = "links-go-tiny-dev"
	}

	return &Config{
		baseUrl:        baseUrl,
		logJson:        logJson,
		linksTableName: linksTableName,
	}
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
