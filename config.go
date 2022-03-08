package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Project   string      `json:"project"`
	Schedules []*Schedule `json:"schedules"`
}

type Schedule struct {
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	AttributionActor string     `json:"attribution-actor"`
	Timetable        Timetable  `json:"timetable"`
	Parameters       Parameters `json:"parameters"`
}

func LoadConfig(configPath string) (*Config, error) {
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var config *Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}
	return config, nil
}
