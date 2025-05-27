package initialize

import (
	"context"
	"fmt"
	"log"

	"github.com/erasernoob/JARVIS/global"
	"github.com/erasernoob/JARVIS/model"
	"github.com/jackc/pgx/v5"
)

func InitPostgresDB(ctx context.Context, config *model.PgDbConfig) (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.Username, config.Password, config.Host, config.Port, config.Database)
	fmt.Println(config)
	fmt.Println(connString)

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	global.PgConn = conn
	if global.PgConn == nil {
		log.Fatalf("Failed to initialize PostgreSQL connection")
		return nil, fmt.Errorf("failed to initialize PostgreSQL connection")
	}
	return conn, nil
}
