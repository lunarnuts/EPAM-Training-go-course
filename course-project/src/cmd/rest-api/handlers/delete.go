package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/src/db/models"
)

func Delete(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil { // bad request
		w.WriteHeader(400)
		return
	}
	err = records.Delete(p, id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}
