package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Browser struct {
		Headless  bool   `yaml:"headless"`
		UserAgent string `yaml:"user_agent"`
		MinWidth  int    `yaml:"min_width"`
		MaxWidth  int    `yaml:"max_width"`
		MinHeight int    `yaml:"min_height"`
		MaxHeight int    `yaml:"max_height"`
	} `yaml:"browser"`

	Timing struct {
		MinDelayMs int `yaml:"min_delay_ms"`
		MaxDelayMs int `yaml:"max_delay_ms"`
	} `yaml:"timing"`

	Limits struct {
		DailyConnections int `yaml:"daily_connections"`
	} `yaml:"limits"`

	Schedule struct {
		StartHour int `yaml:"start_hour"`
		EndHour   int `yaml:"end_hour"`
	} `yaml:"schedule"`
}

func Load(path string) *Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	validate(&cfg)
	return &cfg
}

func validate(cfg *Config) {
	if cfg.Timing.MinDelayMs > cfg.Timing.MaxDelayMs {
		log.Fatal("min_delay_ms cannot exceed max_delay_ms")
	}
	if cfg.Limits.DailyConnections <= 0 {
		log.Fatal("daily_connections must be > 0")
	}
}
