package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/cmd/rest-api/lib"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/db/models"
)

func SelectAll(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	recs, err := records.SelectAll(p)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	lib.ReturnJSON(w, recs)
}
