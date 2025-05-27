package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/erasernoob/JARVIS/src/beans"
)

// 配置文件filepath
var (
	filepath string
	config   *beans.Config
	once     sync.Once
)

// 切换工作目录到项目的根路径
func CheckTheWd() {
	cur, _ := os.Getwd()
	for strings.Contains(cur, "src") {
		_ = os.Chdir("..")
		cur, _ = os.Getwd()
	}
	fmt.Println(cur)
}

func ReadPgDbConfig() (*beans.PgDbConfig, error) {
	return GetConfig().PostgresDbConfig, nil
}

func GetConfig() *beans.Config {
	once.Do(func() {
		var err error
		config = &beans.Config{
			PostgresDbConfig: new(beans.PgDbConfig),
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
