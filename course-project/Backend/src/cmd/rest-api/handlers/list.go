package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/cmd/rest-api/lib"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/db/db"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/db/models"
)

func SelectAll(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := db.AcquireConn(p)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	recs, err := records.SelectAll(conn)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	if len(recs) < 1 {
		recs = make([]records.Record, 0)
	}
	lib.ReturnJSON(w, recs)
}
