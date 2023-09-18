package util

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var maxTries = 10

func NewPgxPool(dbUrl string, ctx context.Context) (pool *pgxpool.Pool, err error) {
	pool, err = pgxpool.New(ctx, dbUrl)
	if err != nil {
		return
	}

	for i := 0; i < maxTries; i++ {
		err = pool.Ping(ctx)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
		continue
	}

	return
}

//func NewSqlDatabase(dbUrl string) (pool *sql.DB, err error) {
//	maxTries := 10
//
//	pool, err = sql.Open("pgx", dbUrl)
//	if err != nil {
//		return
//	}
//
//	for i := 0; i < maxTries; i++ {
//		err = pool.Ping()
//		if err == nil {
//			break
//		}
//		time.Sleep(1 * time.Second)
//		continue
//	}
//
//	return
//}
