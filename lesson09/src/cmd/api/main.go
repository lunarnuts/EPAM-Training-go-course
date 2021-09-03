package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/lunarnuts/go-course/tree/lesson09/src/cmd/api/handlers"
	"github.com/lunarnuts/go-course/tree/lesson09/src/cmd/migrate"
	"github.com/lunarnuts/go-course/tree/lesson09/src/db"
	utils "github.com/lunarnuts/go-course/tree/lesson09/src/env"
)

func dbsFromEnv() db.DBSetup {
	dbs := db.DBSetup{
		User:   utils.EnvOrDef("DB_USER", "postgres"),
		Passwd: utils.EnvOrDef("DB_PASSWD", "postgres"),
		Host:   utils.EnvOrDef("DB_HOST", "localhost"),
		Port:   utils.EnvOrDefInt("DB_PORT", 5432),
		Name:   utils.EnvOrDef("DB_NAME", "postgres"),
		Type:   "postgres",
	}
	return dbs
}

func main() {
	r := mux.NewRouter()
	pool, err := db.New(dbsFromEnv())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()
	err = migrate.MigrateUp(dbsFromEnv())
	if err != nil {
		log.Print(err)
	}
	r.HandleFunc("/api/v1/contacts",
		func(w http.ResponseWriter, r *http.Request) {
			conn, _ := db.AcquireConn(pool)
			handlers.List(&conn, w, r)
		}).Methods("GET")

	r.HandleFunc("/api/v1/contacts/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			conn, _ := db.AcquireConn(pool)
			handlers.Read(&conn, w, r)
		}).Methods("GET")

	r.HandleFunc("/api/v1/contacts",
		func(w http.ResponseWriter, r *http.Request) {
			conn, _ := db.AcquireConn(pool)
			handlers.Insert(&conn, w, r)
		}).Methods("POST")

	r.HandleFunc("/api/v1/contacts/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			conn, _ := db.AcquireConn(pool)
			handlers.Update(&conn, w, r)
		}).Methods("PUT")

	r.HandleFunc("/api/v1/contacts/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			conn, _ := db.AcquireConn(pool)
			handlers.Delete(&conn, w, r)
		}).Methods("DELETE")
	r.Path("/api/v1/contacts/assigngroup").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			conn, _ := db.AcquireConn(pool)
			handlers.AssignContactToGroup(&conn, w, r)
		}).Name("AssignContactsToGroup").Methods(http.MethodPut)
	log.Println("Starting API server on 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func OnErrPanic(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}
