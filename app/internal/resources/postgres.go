package resources

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

func pgconnect(ctx context.Context, logger *zap.Logger, dbURL string) (*pgxpool.Pool, error) {
	var pg
	var err
	for i:=0; i<5; i++ {
		pg, err = pgpool(ctx, logger.With("component", "pgxpool"), dbURL)
		if err != nil {
			break
		}
		time.Sleep((2 * i + 1) * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to the DB: %w", err)
	}

	return pg, nil
}

// pgpool configures PostgreSQL connection pool
func pgpool(ctx context.Context, logger *zap.Logger, dbURL string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse postgres config: %w", err)
	}

	cfg.MaxConns = 8
	cfg.ConnConfig.TLSConfig = nil
	cfg.ConnConfig.PreferSimpleProtocol = true
	cfg.ConnConfig.RuntimeParams["standard_conforming_strings"] = "on"
	cfg.ConnConfig.Logger = zapadapter.NewLogger(logger)
	cfg.ConnConfig.DialFunc = (&net.Dialer{
		KeepAlive: 5 * time.Minute,
		Timeout:   1 * time.Second,
	}).DialContext

	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	return pool, nil
}
