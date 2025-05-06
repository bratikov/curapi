package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ConfigError = errors.New("Config error")
	ConfigFile  string
)

type Log struct {
	Level   string `description:"Log level set to proxy logs." json:"level,omitempty" toml:"level,omitempty" yaml:"level,omitempty" export:"true"`
	Format  string `description:"Proxy log format: json | common" json:"format,omitempty" toml:"format,omitempty" yaml:"format,omitempty" export:"true"`
	NoColor bool   `description:"When using the 'common' format, disables the colorized output." json:"noColor,omitempty" toml:"noColor,omitempty" yaml:"noColor,omitempty" export:"true"`

	FilePath   string `description:"Proxy log file path. Stdout is used when omitted or empty." json:"filePath,omitempty" toml:"filePath,omitempty" yaml:"filePath,omitempty"`
	MaxSize    int    `description:"Maximum size in megabytes of the log file before it gets rotated." json:"maxSize,omitempty" toml:"maxSize,omitempty" yaml:"maxSize,omitempty" export:"true"`
	MaxAge     int    `description:"Maximum number of days to retain old log files based on the timestamp encoded in their filename." json:"maxAge,omitempty" toml:"maxAge,omitempty" yaml:"maxAge,omitempty" export:"true"`
	MaxBackups int    `description:"Maximum number of old log files to retain." json:"maxBackups,omitempty" toml:"maxBackups,omitempty" yaml:"maxBackups,omitempty" export:"true"`
	Compress   bool   `description:"Determines if the rotated log files should be compressed using gzip." json:"compress,omitempty" toml:"compress,omitempty" yaml:"compress,omitempty" export:"true"`
}

func SaveConfig(config any, configPath string) error {
	jsonData, err := json.MarshalIndent(config, "", "   ")
	if err != nil {
		return fmt.Errorf("сant serialize config to JSON: %v", err)
	}
	err = os.WriteFile(configPath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("сant write config file to %v: %v", configPath, err)
	}
	return nil
}

func LoadFromFile(config any, configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		} else {
			return fmt.Errorf("сan not open config file: %v", err)
		}
	}
	defer file.Close()
	loadedConfigBinary, _ := io.ReadAll(file)

	err = json.Unmarshal(loadedConfigBinary, config)
	if err != nil {
		return fmt.Errorf("bad JSON file structure: %v", err)
	}
	return nil
}
