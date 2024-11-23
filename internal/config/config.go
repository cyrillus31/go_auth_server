package config

import (
	"flag"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"./storage/sso.db"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

// Must used when function is not going to return err but will panic
func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("Config path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Config file does not exist: " + path)
	}

	var cfg Config
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	yamlDecoder := yaml.NewDecoder(file)
	if err = yamlDecoder.Decode(&cfg); err != nil {
		panic("Failed to decode the config: " + err.Error())
	}
	return &cfg
}

// fetchconfgiPath fetches config path from command line flag or environment variable
// Priority: flag > env > default
// Defaultl values is empty string
func fetchConfigPath() string {
	var result string

	// --config="paht/to/config.yaml"
	flag.StringVar(&result, "config", "", "path to config file")
	flag.Parse()

	if result == "" {
		result = os.Getenv("CONFIG_PATH")
	}

	return result
}
