package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/cmd/rest-api/lib"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/db/models"
)

func Select(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	id, err := lib.IDFromVars(r)
	if err != nil { // bad request
		lib.ReturnClientError(w, err.Error())
		return
	}
	rec, err := records.Select(p, id)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	lib.ReturnJSON(w, rec)
}
