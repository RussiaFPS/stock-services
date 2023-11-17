package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type Db interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

func NewPostgres(ctx context.Context) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	pool, err := pgxpool.New(ctx, getConfigurationPostgres())
	if err != nil {
		log.Fatal("Error, connect to postgres: ", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal("Error, ping postgres: ", err)
	}

	return pool
}
