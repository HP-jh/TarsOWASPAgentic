package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Endpoint string `json:"endpoint"`
	APIKey   string `json:"api_key"`
	Mode     string `json:"mode"`
	Timeout  int    `json:"timeout_seconds"`
}

func LoadConfigFromFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置失败: %w", err)
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}
	if cfg.Endpoint == "" {
		cfg.Endpoint = "https://api.example.com"
	}
	if cfg.Mode == "" {
		cfg.Mode = "stdio"
	}
	if cfg.Timeout == 0 {
		cfg.Timeout = 30
	}
	return &cfg, nil
}
