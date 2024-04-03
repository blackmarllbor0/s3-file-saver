package postgres

import (
	"app/internal/cfg"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type PgxPool struct {
	db  *pgxpool.Config
	ctx context.Context
}

func NewPgConnection(ctx context.Context, cfg cfg.ConfigService) (*PgxPool, error) {
	pgCfg := cfg.GetAppConfig().Repo.Postgres

	db, err := pgxpool.ParseConfig(pgCfg.ConnString)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to create a cfg: %v", err)
	}

	db.MaxConns = int32(pgCfg.MaxCons)
	db.MinConns = int32(pgCfg.MinCons)
	db.MaxConnLifetime = time.Minute * time.Duration(pgCfg.MaxConnLifetimeInMinutes)
	db.MaxConnIdleTime = time.Minute * time.Duration(pgCfg.MaxConnIdleTimeInMinutes)
	db.HealthCheckPeriod = time.Second * time.Duration(pgCfg.HealthCheckPeriodInSeconds)
	db.ConnConfig.ConnectTimeout = time.Second * time.Duration(pgCfg.ConnectTimeoutInSeconds)

	return &PgxPool{ctx: ctx, db: db}, nil
}

func (p PgxPool) GetConnection() (*pgxpool.Conn, error) {
	pool, err := pgxpool.NewWithConfig(p.ctx, p.db)
	if err != nil {
		return nil, fmt.Errorf("postgres: err while creating conn to the db: %v", err)
	}

	conn, err := pool.Acquire(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("postgres: err while acquiring conn from the db pool: %v\", err")
	}

	if err := conn.Ping(p.ctx); err != nil {
		return nil, fmt.Errorf("postgres: could not ping conn")
	}

	return conn, nil
}
