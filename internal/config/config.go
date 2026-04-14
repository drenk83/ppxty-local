package config

import (
	"errors"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey   string
	LMURL    string
	LMModel  string
	Searcher string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Warn(".env file not found")
	}

	c := &Config{}

	c.APIKey = os.Getenv("SERPER_API_KEY")
	if c.APIKey == "" {
		return nil, errors.New("SERPER_API_KEY is not set")
	}

	c.LMURL = os.Getenv("LM_STUDIO_URL")
	if c.LMURL == "" {
		c.LMURL = "http://localhost:1234"
		slog.Debug("LM_STUDIO_URL not set, using default", "url", c.LMURL)
	}

	c.LMModel = os.Getenv("LM_STUDIO_MODEL")
	if c.LMModel == "" {
		c.LMModel = "local-model"
		slog.Debug("LM_STUDIO_MODEL not set, using default", "model", c.LMModel)
	}

	c.Searcher = os.Getenv("SEARCH_PROVIDER")
	if c.Searcher == "" {
		c.Searcher = "serper"
		slog.Debug("SEARCH_PROVIDER not set, using default", "provider", c.Searcher)
	}

	slog.Info("config successfully loaded", "SERPER_API_KEY", c.APIKey[:6]+"...")
	return c, nil
}
