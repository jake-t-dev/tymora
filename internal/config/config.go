package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string `json:"token"`
}

func NewConfig(ctx context.Context) (context.Context, error) {
	_ = godotenv.Load()

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		return ctx, fmt.Errorf("DISCORD_TOKEN environment variable is not set")
	}

	cfg := &Config{
		Token: token,
	}

	ctx = context.WithValue(ctx, "config", cfg)
	return ctx, nil
}
