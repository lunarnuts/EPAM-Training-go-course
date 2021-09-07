package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/cmd/rest-api/lib"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/db/db"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/db/models"
)

func Insert(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := db.AcquireConn(p)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	var rec records.Record
	err = json.NewDecoder(r.Body).Decode(&rec)
	if err != nil { // bad request
		lib.ReturnClientError(w, err.Error())
		return
	}

	id, err := records.Insert(conn, rec)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}

	resp := make(map[string]string, 1)
	resp["id"] = strconv.FormatUint(id, 10)
	lib.ReturnJSON(w, resp)
}
