package common

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

func CreateTx(ctx context.Context, conn *pgxpool.Pool) (pgx.Tx, error, func()) {
	tx, err := conn.Begin(ctx)
	return tx, err, func() {
		if err != nil {
			err := tx.Rollback(ctx)
			if err != nil {
				slog.Error(err.Error())
			}
		} else {
			err := tx.Commit(ctx)
			if err != nil {
				slog.Error(err.Error())
			}
		}
	}
}
