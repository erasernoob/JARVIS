package global

import (
	"context"
	"fmt"
	"log"

	"github.com/erasernoob/JARVIS/src/config"
	"github.com/jackc/pgx/v5"
)

func InitPostgresDB(ctx context.Context, config *config.PgDbConfig) (*pgx.Conn, error) {
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

	return conn, nil
}
