package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Querier interface {
	QueryRow(ctx context.Context, query string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
}

type Adapter struct {
	pool *pgxpool.Pool
}

func New(connStr string) (*Adapter, error) {
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	return &Adapter{pool: pool}, nil

}

func (db *Adapter) DB(ctx context.Context) Querier {
	return db.pool
}

func (db *Adapter) Close() {
	if db.pool != nil {
		db.pool.Close()
	}
}
