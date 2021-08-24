package handlers

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/src/db/models"
)

func SelectAll(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	recs, err := records.SelectAll(p)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(recs)
	if err != nil {
		log.Printf("Unable to encode json: %v", err)
		w.WriteHeader(500)
		return
	}
}
