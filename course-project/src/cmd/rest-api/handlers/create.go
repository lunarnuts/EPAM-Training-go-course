package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/src/db/models"
)

func Insert(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	var rec records.Record
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil { // bad request
		w.WriteHeader(400)
		return
	}

	id, err := records.Insert(p, rec)
	if err != nil {
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
