package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/src/db/db"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/src/db/models/records"
)

func main() {
	router := mux.NewRouter()
	dbs := db.DBSetup{
		User:   "postgres",
		Passwd: "1234",
		Host:   "localhost",
		Port:   5432,
		Name:   "postgres",
		Type:   "postgres",
	}
	pool, err := dbs.New()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/records",
		func(w http.ResponseWriter, r *http.Request) {
			records.SelectAll(pool, w, r)
		}).Methods("GET")

	r.HandleFunc("/api/v1/records/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			records.Select(pool, w, r)
		}).Methods("GET")

	r.HandleFunc("/api/v1/records",
		func(w http.ResponseWriter, r *http.Request) {
			records.Insert(pool, w, r)
		}).Methods("POST")

	r.HandleFunc("/api/v1/records/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			records.Update(pool, w, r)
		}).Methods("PUT")

	r.HandleFunc("/api/v1/records/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			records.Delete(pool, w, r)
		}).Methods("DELETE")
	router.Path("/api/v0/weather").Queries("city", "{cityName}").HandlerFunc().Name("GetCurrentWeather").Methods(http.MethodGet)
	log.Println("Starting API server on 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
