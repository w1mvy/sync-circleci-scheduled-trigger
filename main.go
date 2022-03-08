package main

import (
	"context"
	"flag"
	"fmt"
)

var configPath string

func main() {
	flag.StringVar(&configPath, "config", ".circleci-schedule.json", "path of scheduled trigger json")
	flag.Parse()

	config, err := LoadConfig(configPath)
	if err != nil {
		fmt.Errorf("failed to load config file: %v", err)
	}
	client, err := NewClient()
	if err != nil {
		fmt.Errorf("failed to client: %v", err)
	}
	_, err = Sync(context.Background(), client, config)
	if err != nil {
		fmt.Errorf("failed sync %v", err)
	}
}
