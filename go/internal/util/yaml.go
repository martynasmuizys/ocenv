package util

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type KubeConfig struct {
	ApiVersion        string         `yaml:"apiVersion"`
	Kind              string         `yaml:"kind"`
	Preferences       map[string]any `yaml:"preferences,inline,flow"`
	Clusters          []Clusters     `yaml:"clusters"`
	Users             []Users        `yaml:"users"`
	Contexts          []Contexts     `yaml:"contexts"`
	CurrentContext    string         `yaml:"current-context"`
	OcenvTokenExpires int64          `yaml:"ocenv-token-expires,"`
}

type Clusters struct {
	Cluster Cluster `yaml:"cluster"`
	Name    string  `yaml:"name"`
}

type Cluster struct {
	Server string `yaml:"server"`
}

type Contexts struct {
	Context Context `yaml:"context"`
	Name    string  `yaml:"name"`
}

type Context struct {
	Cluster   string `yaml:"cluster"`
	Namespace string `yaml:"namespace"`
	User      string `yaml:"user"`
}

type Users struct {
	Name string `yaml:"name"`
	User User   `yaml:"user"`
}

type User struct {
	Token string `yaml:"token"`
}

func ParseConfig(path string) (*KubeConfig, error) {
	var cfg KubeConfig

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("Failed to read environment file: %v", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("Failed to parse environment file: %v", err)
	}

	return &cfg, nil
}

func SaveConfig(cfg *KubeConfig, path string) error {
	data, err := yaml.Marshal(cfg)

	if err != nil {
		return fmt.Errorf("Failed to parse config: %v", err)
	}

	// idk why not os.WriteFile
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("Failed to save config: %v", err)
	}

	return nil
}
