package postgres

import (
	"api-std/config"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func PostgresPoolDestroy() {
	PostgresPool.Close()
}

func PostgresPoolInit() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=America/Denver", config.Env.PostgresHost, config.Env.PostgresUser,
		config.Env.PostgresPass, config.Env.PostgresPort)
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		slog.Error("Unable to create connection pool: %v", err)
		return
	}

	// Use a timeout to avoid hanging if the database is down.
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := pool.Ping(pingCtx); err != nil {
		slog.Error("Could not ping database: %v", err)
		return
	}

	slog.Info("PostgresPool has been successfully initialized")
	PostgresPool = pool
}

func PostgresPoolPing() error {
	if PostgresPool == nil {
		return errors.New("PostgresPool is nil")
	}
	ctx := context.Background()

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := PostgresPool.Ping(pingCtx); err != nil {
		slog.Error("Could not ping database: %v", err)
		return err
	}
	return nil
}

var PostgresPool *pgxpool.Pool = nil
