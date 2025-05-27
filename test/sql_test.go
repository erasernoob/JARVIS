package test

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
	"testing"

	"github.com/erasernoob/JARVIS/auth"
	g "github.com/erasernoob/JARVIS/global"
	"github.com/erasernoob/JARVIS/initialize"
	"github.com/jackc/pgx/v5"
)

func InitTestEnv() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("panic occurred: %v", r)
			debug.PrintStack()
		}
	}()
	ctx := context.Background()
	// mock the userID
	ctx = auth.Identify(ctx)

	if err := g.Init(ctx); err != nil {
		log.Fatalf("init failed: %s", err)
	}
	if err := initialize.Init(ctx); err != nil {
		log.Fatalf("initialize failed: %s", err)
	}

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
