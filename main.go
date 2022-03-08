package main

import (
	"context"
	"flag"
	"fmt"
)

var configPath string
var dryRun bool

func main() {
	flag.StringVar(&configPath, "config", ".circleci-schedule.json", "path of scheduled trigger json")
	flag.BoolVar(&dryRun, "dryrun", false, "enabled dry-run mode")
	flag.Parse()

	config, err := LoadConfig(configPath)
	if err != nil {
		fmt.Errorf("failed to load config file: %v", err)
	}
	client, err := NewClient()
	if err != nil {
		fmt.Errorf("failed to client: %v", err)
	}

	if dryRun {
		client.Logger.Println("Execute as dry-run mode")
	}
	_, err = Sync(context.Background(), client, config, dryRun)
	if err != nil {
		fmt.Errorf("failed sync %v", err)
	}
}
