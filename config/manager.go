package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/erasernoob/JARVIS/model"
)

// 配置文件filepath
var (
	filepath string
	config   *model.Config
	once     sync.Once
)

// 切换工作目录到项目的根路径
func CheckTheWd() {
	cur, _ := os.Getwd()
	for strings.Contains(cur, "test") {
		_ = os.Chdir("..")
		cur, _ = os.Getwd()
	}
	fmt.Println(cur)
}

func ReadPgDbConfig() (*model.PgDbConfig, error) {
	return GetConfig().PostgresDbConfig, nil
}

func GetConfig() *model.Config {
	once.Do(func() {
		var err error
		config = &model.Config{
			PostgresDbConfig: new(model.PgDbConfig),
		}
		if err = ReadJsonConfigFile(filepath); err != nil {
			log.Fatalf("Failed to read config file: %v", err)
			return
		}
	})
	return config
}

func init() {
	CheckTheWd()

	flag.StringVar(&filepath, "config", "dev.json", "Path to the configuration file")

	_ = GetConfig()

	// flag.Parse()

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
