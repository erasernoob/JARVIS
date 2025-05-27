package initialize

import (
	"context"

	"github.com/erasernoob/JARVIS/config"
	"github.com/erasernoob/JARVIS/global"
)

func Init(ctx context.Context) error {
	// Init the config
	pgConfig, err := config.ReadPgDbConfig()
	if err != nil {
		// log.Fatalf("Failed to read PostgreSQL config: %v", err)
		return err
	}

	// Initialize PostgreSQL database connection
	global.PgConn, err = InitPostgresDB(ctx, pgConfig)
	if err != nil {
		// log.Fatalf("Failed to initialize PostgreSQL database: %v", err)
		return err
	}

	// Create the Agent
	global.Agent, err = InitAgent(ctx)
	if err != nil {
		// log.Fatalf("Failed to initialize agent: %v", err)
		return err
	}
	return nil

}
