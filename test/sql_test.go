package test

import (
	"context"
	"fmt"
	"testing"

	g "github.com/erasernoob/JARVIS/global"
	"github.com/jackc/pgx/v5"
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

	if err := createTables(g.PgConn); err != nil {
		fmt.Println(err)
	}
}

func createTables(con *pgx.Conn) error {
	ctx := context.Background()

	createConversations := `
	CREATE TABLE IF NOT EXISTS conversations (
		id UUID PRIMARY KEY,
		user_id TEXT,
		title TEXT,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		metadata JSONB
	);`

	createMessages := `
	CREATE TABLE IF NOT EXISTS messages (
		id UUID PRIMARY KEY,
		conversation_id UUID REFERENCES conversations(id) ON DELETE CASCADE,
		role TEXT NOT NULL,
		content TEXT NOT NULL,
		timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		metadata JSONB
	);`

	_, err := con.Exec(ctx, createConversations)
	if err != nil {
		return fmt.Errorf("failed to create conversations table: %w", err)
	}

	_, err = con.Exec(ctx, createMessages)
	if err != nil {
		return fmt.Errorf("failed to create messages table: %w", err)
	}

	return nil
}
