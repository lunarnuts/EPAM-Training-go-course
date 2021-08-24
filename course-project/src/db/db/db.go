package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
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

	pool, err := pgxpool.Connect(context.Background(), dbs.String())
	if err != nil {
		log.Fatalf("Unable to connection to database: %v", err)
	}
	return pool, err
}
