package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Port    string `yaml:"port" required:"true"`
	MongoDB string `yaml:"MongoDB" required:"true"`
}

// MustLoad loads the configuration from the default path.
func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config file path is empty")
	}

	return MustLoadByPath(path)
}

func fetchConfigPath() string {
	var path string

	// --config="/path/to/config.yaml"
	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}

// MustLoadByPath loads the configuration from the specified path.
func MustLoadByPath(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found")
	}
	var config Config

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("failed to read config file")
	}

	return &config
}
