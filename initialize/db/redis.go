package db

import (
	"context"
	"os"

	"github.com/erasernoob/JARVIS/common"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context) (cli redis.Client) {
	cli = *redis.NewClient(&redis.Options{
		Addr:     os.Getenv(common.REDIS_ADDR),
		Protocol: 2,
	})
	return
}
