package main

import (
	"log/slog"
	"os"

	"github.com/drenk83/ppxty-local/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "err", err)
		os.Exit(1)
	}
	_ = cfg
}
