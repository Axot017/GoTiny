package internal

import "os"

type Config struct {
	baseUrl string
}

func NewConfig() *Config {
	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost:8080"
	}
	return &Config{baseUrl: baseUrl}
}

func (c *Config) BaseUrl() string {
	return c.baseUrl
}
