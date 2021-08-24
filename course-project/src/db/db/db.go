package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type DBSetup struct {
	User   string
	Passwd string
	Host   string
	Port   int
	Name   string
	Type   string
}

func (dbs DBSetup) String() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
		dbs.Type, dbs.User, dbs.Passwd, dbs.Host, dbs.Port, dbs.Name)
}

func New(dbs DBSetup) (*pgxpool.Pool, error) {
	t, err := sql.Open(dbs.Type, dbs.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to open DB")
	}
	return t, nil
	pool, err := pgxpool.Connect(context.Background(), dbs.String())
	if err != nil {
		log.Fatalf("Unable to connection to database: %v", err)
	}
	defer pool.Close()
}
