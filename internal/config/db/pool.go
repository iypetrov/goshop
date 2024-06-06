package db

import (
	"context"
	"fmt"
	"github.com/iypetrov/goshop/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDatabaseConnectionPool(ctx context.Context) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(
		ctx,
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=%s",
			config.Get().Storage.Username,
			config.Get().Storage.Password,
			config.Get().Storage.Addr,
			config.Get().Storage.Name,
			config.Get().Storage.SSL,
		),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
