package global

import (
	"context"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
	"github.com/erasernoob/JARVIS/src/agent"
	"github.com/erasernoob/JARVIS/src/beans"
	"github.com/erasernoob/JARVIS/src/config"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var (
	LLM     model.ToolCallingChatModel
	History []*schema.Message
	Agent   *beans.Client
	PgConn  *pgx.Conn
)

func Init(ctx context.Context) error {
	// Load environment variables from .env file
	var err error
	if err := godotenv.Load(); err != nil {
		return err
	}

	// Init the config
	pgConfig, err := config.ReadPgDbConfig()
	if err != nil {
		// log.Fatalf("Failed to read PostgreSQL config: %v", err)
		return err
	}

	// Initialize PostgreSQL database connection
	PgConn, err = InitPostgresDB(ctx, pgConfig)
	if err != nil {
		// log.Fatalf("Failed to initialize PostgreSQL database: %v", err)
		return err
	}

	// Create the Agent
	Agent, err = agent.InitAgent(ctx)
	if err != nil {
		// log.Fatalf("Failed to initialize agent: %v", err)
		return err
	}

	return nil
}
