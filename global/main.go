package global

import (
	"context"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"

	m "github.com/erasernoob/JARVIS/model"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var (
	LLM     model.ToolCallingChatModel
	History []*schema.Message
	Agent   *m.Client
	PgConn  *pgx.Conn
)

func Init(ctx context.Context) error {
	// Load environment variables from .env file
	// var err error
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}
