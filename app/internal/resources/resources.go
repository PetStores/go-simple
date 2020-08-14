package resources

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type R struct {
	Config Config
	PGPool *pgxpool.Pool
}

type Config struct {
	DiagPort    int    `envconfig:"DIAG_PORT" default:"8081" required:"true"`
	RESTAPIPort int    `envconfig:"PORT" default:"8080" required:"true"`
	DBURL       string `envconfig:"DATABASE_URL" default:"postgres://user:password@localhost:5432/petstore?sslmode=disable" required:"true"`
}

func New(ctx context.Context, logger *zap.Logger) (*R, error) {
	conf := Config{}
	err := envconfig.Process("", &conf)
	if err != nil {
		return nil, fmt.Errorf("can't process the config: %w", err)
	}

	pgpool, err := pgconnect(ctx, logger, conf.DBURL)
	if err != nil {
		return nil, err
	}

	return &R{
		Config: conf,
		PGPool: pgpool,
	}, nil
}
