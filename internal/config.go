package internal

import (
	"os"
	"strconv"
)

type Config struct {
	baseUrl         string
	logJson         bool
	linksTableName  string
	maxTrackingDays uint
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
	maxTrackingDaysStr := os.Getenv("MAX_TRACKING_DAYS")
	maxTrackingDays, err := strconv.Atoi(maxTrackingDaysStr)
	if err != nil {
		maxTrackingDays = 30
	}

	return &Config{
		baseUrl:         baseUrl,
		logJson:         logJson,
		linksTableName:  linksTableName,
		maxTrackingDays: uint(maxTrackingDays),
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

func (c *Config) MaxTrackingDays() uint {
	return c.maxTrackingDays
}
