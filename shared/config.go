package shared

import (
	"log"
	"net/http"
)

type Config struct {
	BaseURL    string
	UserAgent  string
	HttpClient *http.Client
	Logger     func(string, ...any)
	Debug      bool
	CompanyID  *string
}

func NewConfig() *Config {
	return &Config{
		BaseURL:    "https://external.pleo.io",
		Logger:     log.Printf,
		Debug:      false,
		HttpClient: http.DefaultClient,
		UserAgent:  "Die-Agenturverwaltung/1.0.0/go",
	}
}

func NewStagingConfig() *Config {
	return &Config{
		BaseURL:    "https://external.staging.pleo.io",
		Logger:     log.Printf,
		Debug:      false,
		HttpClient: http.DefaultClient,
		UserAgent:  "Die-Agenturverwaltung/1.0.0/go",
	}
}
