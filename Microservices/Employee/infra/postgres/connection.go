package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/rhzyyn/Aozora/Microservices/Employee/config"
	"github.com/rhzyyn/Aozora/Microservices/Employee/infra/logger"

	"github.com/cenkalti/backoff/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type DB struct {
	Pool   *pgxpool.Pool
	Tracer trace.Tracer
}

func Connect(ctx context.Context, cfg *config.Config, log *logger.Logger) (*DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_max_conns=10 pool_max_conn_lifetime=5m",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
		cfg.DB.SSLMode,
	)

	var dbpool *pgxpool.Pool
	operation := func() error {
		var err error
		dbpool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			log.Warn("retrying DB connection", "err", err)
			return err
		}
		return dbpool.Ping(ctx)
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 30 * time.Second

	if err := backoff.Retry(operation, backoff.WithContext(b, ctx)); err != nil {
		log.Error("failed to connect to PostgreSQL", "dsn", dsn, "err", err)
		return nil, err
	}

	log.Info("PostgreSQL connected successfully")

	return &DB{
		Pool:   dbpool,
		Tracer: otel.Tracer("postgres"),
	}, nil
}

func (db *DB) Close() {
	db.Pool.Close()
}
