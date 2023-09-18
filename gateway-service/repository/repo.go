package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	//_ "github.com/jackc/pgx/v5/stdlib"
)

type Repo Querier

func NewSqlRepository(db *pgxpool.Pool) Repo {
	return New(db)
}
