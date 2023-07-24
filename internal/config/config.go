package config

import (
	"os"
)

type Config struct {
	baseUrl        string
	logJson        bool
	linksTableName string
	ipTableName    string
	ipStackBaseUrl string
	ipStackToken   string
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
	ipTableName := os.Getenv("IP_DYNAMODB_TABLE")
	if ipTableName == "" {
		ipTableName = linksTableName
	}

	ipStackBaseUrl := os.Getenv("IP_STACK_BASE_URL")
	if ipStackBaseUrl == "" {
		ipStackBaseUrl = "http://api.ipstack.com"
	}
	ipStackToken := os.Getenv("IP_STACK_TOKEN")

	return &Config{
		baseUrl:        baseUrl,
		logJson:        logJson,
		linksTableName: linksTableName,
		ipTableName:    ipTableName,
		ipStackBaseUrl: ipStackBaseUrl,
		ipStackToken:   ipStackToken,
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

func (c *Config) IpTableName() string {
	return c.ipTableName
}

func (c *Config) IpStackBaseUrl() string {
	return c.ipStackBaseUrl
}

func (c *Config) IpStackToken() string {
	return c.ipStackToken
}
