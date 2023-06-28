package internal

import "os"

type Config struct {
	baseUrl string
	logJson bool
}

func NewConfig() *Config {
	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost:8080"
	}
	logJson := os.Getenv("LOG_JSON")

	return &Config{baseUrl: baseUrl, logJson: logJson == "true"}
}

func (c *Config) BaseUrl() string {
	return c.baseUrl
}

func (c *Config) LogJson() bool {
	return c.logJson
}
