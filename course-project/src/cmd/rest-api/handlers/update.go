package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/src/db/models"
)

func Update(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil { // bad request
		w.WriteHeader(400)
		return
	}

	var rec records.Record
	err = json.NewDecoder(r.Body).Decode(&rec)
	if err != nil { // bad request
		w.WriteHeader(400)
		return
	}
	err = records.Update(p, id, rec)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}
