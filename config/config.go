package config

import (
	"github.com/pelletier/go-toml/v2"
	"io/ioutil"
	"serversdf/server"
)

type Config struct {
	KnownHosts string              `toml:"known_hosts"`
	Servers    []server.ServerInfo `toml:"server"`
}

func ReadConfigurationFile(file string) (*Config, error) {
	var cfg Config
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(content, &cfg)
	return &cfg, err
}

func (cfg *Config) GetNumberOfServers() int {
	return len(cfg.Servers)
}

func (cfg *Config) GetLongestServerName() int {
	longestServerName := 0

	for _, serverInfo := range cfg.Servers {
		currentServerNameLength := len(serverInfo.Name)
		if currentServerNameLength > longestServerName {
			longestServerName = currentServerNameLength
		}
	}

	return longestServerName
}
