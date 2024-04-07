package postgres

import (
	"app/internal/cfg"
	"app/pkg/logger"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type PgxPool struct {
	db *pgxpool.Config

	ctx context.Context

	logService logger.LoggerService
}

func NewPgConnection(
	ctx context.Context,
	cfg cfg.ConfigService,
	logService logger.LoggerService,
) (*PgxPool, error) {
	pgCfg := cfg.GetAppConfig().Repo.Postgres

	db, err := pgxpool.ParseConfig(pgCfg.ConnString)
	if err != nil {
		err := fmt.Errorf("postgres.NewPgConnection: failed to create a cfg: %v", err)

		logService.Error(err.Error())

		return nil, err
	}

	db.MaxConns = int32(pgCfg.MaxCons)
	db.MinConns = int32(pgCfg.MinCons)
	db.MaxConnLifetime = time.Minute * time.Duration(pgCfg.MaxConnLifetimeInMinutes)
	db.MaxConnIdleTime = time.Minute * time.Duration(pgCfg.MaxConnIdleTimeInMinutes)
	db.HealthCheckPeriod = time.Second * time.Duration(pgCfg.HealthCheckPeriodInSeconds)
	db.ConnConfig.ConnectTimeout = time.Second * time.Duration(pgCfg.ConnectTimeoutInSeconds)

	go logService.Info("postgres.NewPgConnection: successfully configuration postgres")

	return &PgxPool{
		ctx:        ctx,
		db:         db,
		logService: logService,
	}, nil
}

func (p PgxPool) GetConnection() (*pgxpool.Conn, error) {
	pool, err := pgxpool.NewWithConfig(p.ctx, p.db)
	if err != nil {
		err := fmt.Errorf("postgres.GetConnection: err while creating conn to the db: %v", err)

		p.logService.Error(err.Error())

		return nil, err
	}

	go p.logService.Info("postgres.GetConnection: successfully created pool")

	conn, err := pool.Acquire(p.ctx)
	if err != nil {
		err := fmt.Errorf("postgres.GetConnection: err while acquiring conn from the db pool: %v", err)

		p.logService.Error(err.Error())

		return nil, err
	}

	go p.logService.Info("postgres.GetConnection: successfully get connection from pool")

	if err := conn.Ping(p.ctx); err != nil {
		err := fmt.Errorf("postgres.GetConnection: could not ping conn")

		p.logService.Error(err.Error())

		return nil, err
	}

	go p.logService.Info("postgres.GetConnection: connection has been verified")

	return conn, nil
}
