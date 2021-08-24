package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/src/cmd/rest-api/handlers"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/src/db/db"
)

func main() {
	r := mux.NewRouter()
	dbs := db.DBSetup{
		User:   "postgres",
		Passwd: "1234",
		Host:   "localhost",
		Port:   5432,
		Name:   "postgres",
		Type:   "postgres",
	}
	pool, err := db.New(dbs)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	r.HandleFunc("/api/v1/records",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.SelectAll(pool, w, r)
		}).Methods("GET")

	r.HandleFunc("/api/v1/records/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.Select(pool, w, r)
		}).Methods("GET")

	r.HandleFunc("/api/v1/records",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.Insert(pool, w, r)
		}).Methods("POST")

	r.HandleFunc("/api/v1/records/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.Update(pool, w, r)
		}).Methods("PUT")

	r.HandleFunc("/api/v1/records/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.Delete(pool, w, r)
		}).Methods("DELETE")
	r.Path("/api/v1/weather").Queries("city", "{cityName}").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handlers.GetCurrentWeather(pool, w, r)
		}).Name("GetCurrentWeather").Methods(http.MethodGet)
	log.Println("Starting API server on 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
