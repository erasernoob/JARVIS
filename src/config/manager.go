package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sync"
)

// 配置文件filepath
var (
	filepath string
	config   *Config
	once     sync.Once
)

func ReadPgDbConfig() (*PgDbConfig, error) {
	return GetConfig().PostgresDbConfig, nil
}

func GetConfig() *Config {
	once.Do(func() {
		var err error
		config = &Config{
			PostgresDbConfig: new(PgDbConfig),
		}
		if err = ReadJsonConfigFile(filepath); err != nil {
			log.Fatalf("Failed to read config file: %v", err)
			os.Exit(1)
		}
	})
	return config
}

func init() {
	flag.StringVar(&filepath, "config", "dev.json", "Path to the configuration file")

	_ = GetConfig()

	flag.Parse()

}

func ReadJsonConfigFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}
	return nil
}
