package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/lunarnuts/go-course/tree/lesson09/src/db"
	utils "github.com/lunarnuts/go-course/tree/lesson09/src/env"
	"github.com/lunarnuts/go-course/tree/lesson09/src/models"
)

func dbsFromEnv() db.DBSetup {
	dbs := db.DBSetup{
		User:   utils.EnvOrDef("DB_USER", "postgres"),
		Passwd: utils.EnvOrDef("DB_PASSWD", "1234"),
		Host:   utils.EnvOrDef("DB_HOST", "localhost"),
		Port:   utils.EnvOrDefInt("DB_PORT", 5432),
		Name:   utils.EnvOrDef("DB_NAME", "course_db"),
		Type:   "postgres",
	}
	return dbs
}

func main() {
	dbObj, err := db.New(dbsFromEnv())
	OnErrPanic(err)
	users, err := models.ContactList(dbObj)
	OnErrPanic(err)
	fmt.Println(users)
}

func OnErrPanic(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}
