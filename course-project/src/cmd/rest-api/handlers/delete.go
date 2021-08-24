package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/src/cmd/rest-api/lib"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/src/db/models"
)

func Delete(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	id, err := lib.IDFromVars(r)
	if err != nil { // bad request
		lib.ReturnClientError(w, err.Error())
		return
	}
	err = records.Delete(p, id)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	w.WriteHeader(200)
}
