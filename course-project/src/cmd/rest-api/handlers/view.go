package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"log"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/src/db/models"
)

func Select(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil { // bad request
		w.WriteHeader(400)
		return
	}
	rec, err := records.Select(p, id)
	if err != nil {
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
