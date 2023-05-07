package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Pane struct {
	Command string `yaml:"command"` // the command to run
	Delay   int64  `yaml:"delay"`   // how long to wait before running the command
}

type Config struct {
	Name  string `yaml:"name"`
	Panes []Pane `yaml:"panes"`
}

func LoadConfig(file string) (*Config, error) {
	file, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}
	if !fileExists(file) {
		return nil, fmt.Errorf("file does not exist")
	}
	c := &Config{
		Name:  "",
		Panes: []Pane{},
	}
	b, err := os.ReadFile(file)
	yaml.Unmarshal(b, c)
	return c, err
}

func fileExists(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()
	_, err = file.Stat()
	return err == nil
}
