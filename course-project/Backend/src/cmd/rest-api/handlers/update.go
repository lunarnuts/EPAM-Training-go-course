package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/src/cmd/rest-api/lib"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/src/db/models"
)

func Update(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	id, err := lib.IDFromVars(r)
	if err != nil { // bad request
		lib.ReturnClientError(w, err.Error())
		return
	}

	var rec records.Record
	err = json.NewDecoder(r.Body).Decode(&rec)
	if err != nil { // bad request
		lib.ReturnClientError(w, err.Error())
		return
	}
	err = records.Update(p, id, rec)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}

	w.WriteHeader(200)
}
