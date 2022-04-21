package main

import (
	"context"
	"flag"
	"fmt"
	"os"
)

var configPath string
var dryRun bool
var forceSync bool

func main() {
	flag.StringVar(&configPath, "config", ".circleci-schedule.json", "path of scheduled trigger json")
	flag.BoolVar(&dryRun, "dryrun", false, "enabled dry-run mode")
	flag.BoolVar(&forceSync, "forcesync", false, "delete schedules that does not exist in config. judge by only name match.")
	flag.Parse()

	config, err := LoadConfig(configPath)
	if err != nil {
		fmt.Printf("failed to load config file: %v\n", err)
		os.Exit(1)
	}
	client, err := NewClient()
	if err != nil {
		fmt.Printf("failed to create client: %v\n", err)
		os.Exit(1)
	}

	if dryRun {
		fmt.Println("Execute as dry-run mode")
	}
	if forceSync {
		fmt.Println("Execute as force-sync mode")
	}
	_, err = Sync(context.Background(), client, config, dryRun, forceSync)
	if err != nil {
		fmt.Printf("failed sync %v\n", err)
		os.Exit(1)
	}
}
