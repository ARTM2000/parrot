package core

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"slices"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type config struct {
	Version  int             `yaml:"version"`
	Server   serverConfig    `yaml:"server"`
	Requests []requestConfig `yaml:"requests"`
}

type serverConfig struct {
	Port  int  `yaml:"port"`
	Watch bool `yaml:"watch"`
}

type requestConfig struct {
	Method       string `yaml:"method"`
	Path         string `yaml:"path"`
	ResponseFile string `yaml:"response_file"`
}

func LoadConfig(configPath string) *config {
	c := readConfigFromFile(configPath)

	if c.Server.Watch {
		go func() { c.watchConfigFile(configPath) }()
	}

	return c
}

func readConfigFromFile(configPath string) *config {
	file, err := os.ReadFile(configPath)
	if os.IsNotExist(err) {
		slog.Error("config file not found")
		return nil
	}

	var c config
	if err = yaml.Unmarshal(file, &c); err != nil {
		slog.Error("fail to parse config", slog.Any("error", err))
		return nil
	}

	return &c
}

func (c *config) IsValid() bool {
	if c.Version <= 0 {
		slog.Error("config version is not valid. should be bigger than 0")
		return false
	}

	if c.Server.Port == 0 {
		slog.Error("server.port is invalid. port should be a valid port number")
		return false
	}

	for _, r := range c.Requests {
		if !slices.Contains([]string{"GET", "POST", "PATCH", "DELETE", "PUT"}, r.Method) {
			slog.Error(fmt.Sprintf("invalid request method: `%s`\n", r.Method))
			return false
		}

		if strings.Contains(r.ResponseFile, "${PWD}") {
			dir, _ := os.Getwd()
			r.ResponseFile = strings.Replace(r.ResponseFile, "${PWD}", dir, 1)
		}

		if _, err := os.Stat(r.ResponseFile); os.IsNotExist(err) {
			slog.Error(fmt.Sprintf("could not resolve `%s` response file\n", r.ResponseFile))
			return false
		}

		if _, err := regexp.Compile(r.Path); err != nil {
			slog.Error(fmt.Sprintf("unable to parse request path `%s`", r.Path))
			return false
		}
	}

	return true
}

func (c *config) watchConfigFile(path string) error {
	lastStats, _ := os.Stat(path)

	for {
		stat, err := os.Stat(path)
		if err != nil {
			return err
		}

		if stat.Size() != lastStats.Size() || stat.ModTime() != lastStats.ModTime() {
			lastStats = stat
			newConfig := readConfigFromFile(path)
			if newConfig.IsValid() {
				c = newConfig // to hot reload the config file
				slog.Info("[Hot Reload] config file updated")
			} else {
				slog.Error("[Hot Reload] config file is invalid")
			}
		}

		time.Sleep(2 * time.Second)
	}
}
