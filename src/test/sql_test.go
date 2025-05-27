package test

import (
	"context"
	"fmt"
	"testing"

	g "github.com/erasernoob/JARVIS/src/global"
	"github.com/joho/godotenv"
)

var ctx context.Context

func InitTestEnv() {
	_ = godotenv.Load()
	ctx = context.Background()
	_ = g.Init(ctx) // 你原来的 g.Init(ctx)
}

func Test_SQL(t *testing.T) {
	InitTestEnv()

	tag, err := g.PgConn.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS conversations (
    id UUID PRIMARY KEY,
    user_id TEXT,
    title TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    metadata JSONB
);`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tag)
}
