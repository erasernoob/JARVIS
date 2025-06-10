package db

import (
	"context"
	"os"

	"github.com/erasernoob/JARVIS/common"
	pkgRedis "github.com/erasernoob/JARVIS/pkg/redis"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context) (cli redis.Client) {
	cli = *redis.NewClient(&redis.Options{
		Addr:     os.Getenv(common.REDIS_ADDR),
		Protocol: 2,
	})
	return
}

func init() {
	// 创建redis索引
	if err := pkgRedis.Init(); err != nil {
		panic("Failed to initialize Redis index: " + err.Error())
	}
}
