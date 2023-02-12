package main

import "path/filepath"

type AppConfig struct {
	configPath string
	storePath  string
}

func NewConfig() *AppConfig {
	return &AppConfig{
		storePath:  "store.json",
		configPath: "~/.config/todo/",
	}
}

func (c *AppConfig) StorePath() string {
	return filepath.Join(c.configPath, c.storePath)
}
