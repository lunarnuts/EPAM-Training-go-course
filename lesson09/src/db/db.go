package db

import (
	"database/sql"
	"fmt"
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
	return fmt.Sprintf("%s://%s:%ss@%s:%d/%s?sslmode=disable",
		dbs.Type, dbs.User, dbs.Passwd, dbs.Host, dbs.Port, dbs.Name)
}

func New(dbs DBSetup) (*sql.DB, error) {
	t, er := sql.Open(dbs.Type, dbs.String())
	if er != nil {
		return nil, er
	}
	return t, nil
}
