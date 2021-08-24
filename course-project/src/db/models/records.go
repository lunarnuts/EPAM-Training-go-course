package records

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"log"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Record struct {
	Id            int    `json:"id", db:"id"`
	CityName      string `json:"cityName", db:"cityName"`
	TimeRequested string `json:"timeRequested", db:"timeRequested"`
	Temperature   string `json:"temperature", db:"temperature"`
}

func SelectAll(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil { // bad request
		w.WriteHeader(400)
		return
	}

	conn, err := p.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		w.WriteHeader(500)
		return
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		"SELECT * FROM logbook",
		id)

	var rec Record
	err = row.Scan(&rec.Id, &rec.CityName, &rec.TimeRequested, &rec.Temperature)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		return
	}

	if err != nil {
		log.Printf("Unable to SELECT: %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(rec)
	if err != nil {
		log.Printf("Unable to encode json: %v", err)
		w.WriteHeader(500)
		return
	}
}

func Select(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil { // bad request
		w.WriteHeader(400)
		return
	}

	conn, err := p.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		w.WriteHeader(500)
		return
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		"SELECT id, citynName, timeRequested, temperature FROM logbook WHERE id = $1",
		id)

	var rec Record
	err = row.Scan(&rec.Id, &rec.CityName, &rec.TimeRequested, &rec.Temperature)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		return
	}

	if err != nil {
		log.Printf("Unable to SELECT: %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(rec)
	if err != nil {
		log.Printf("Unable to encode json: %v", err)
		w.WriteHeader(500)
		return
	}
}

func Insert(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	var rec Record
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil { // bad request
		w.WriteHeader(400)
		return
	}

	conn, err := p.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v", err)
		w.WriteHeader(500)
		return
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		"INSERT INTO logbook (cityName, timeRequested) VALUES ($1, $2) RETURNING id",
		rec.CityName, rec.TimeRequested)
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		log.Printf("Unable to INSERT: %v", err)
		w.WriteHeader(500)
		return
	}

	resp := make(map[string]string, 1)
	resp["id"] = strconv.FormatUint(id, 10)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("Unable to encode json: %v", err)
		w.WriteHeader(500)
		return
	}
}

func Update(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil { // bad request
		w.WriteHeader(400)
		return
	}

	var rec Record
	err = json.NewDecoder(r.Body).Decode(&rec)
	if err != nil { // bad request
		w.WriteHeader(400)
		return
	}

	conn, err := p.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v", err)
		w.WriteHeader(500)
		return
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(),
		"UPDATE logbook SET cityName = $2, timeRequested = $3, temperature = $4 WHERE id = $1",
		id, rec.CityName, rec.TimeRequested, rec.Temperature)
	if err != nil {
		log.Printf("Unable to UPDATE: %v\n", err)
		w.WriteHeader(500)
		return
	}

	if ct.RowsAffected() == 0 {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
}

func Delete(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil { // bad request
		w.WriteHeader(400)
		return
	}

	conn, err := p.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v", err)
		w.WriteHeader(500)
		return
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(), "DELETE FROM logbook WHERE id = $1", id)
	if err != nil {
		log.Printf("Unable to DELETE: %v", err)
		w.WriteHeader(500)
		return
	}

	if ct.RowsAffected() == 0 {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
}
