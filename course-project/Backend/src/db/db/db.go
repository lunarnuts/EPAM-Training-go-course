package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
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

type Pooler interface {
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
}

type DBConn interface {
	//Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
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

func AcquireConn(p Pooler) (conn DBConn, err error) {
	dbConn, err := p.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return conn, err
	}
	conn = dbConn
	return conn, nil
}
