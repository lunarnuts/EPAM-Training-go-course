package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/src/cmd/rest-api/lib"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/src/db/models"
)

func Insert(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	var rec records.Record
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil { // bad request
		lib.ReturnClientError(w, err.Error())
		return
	}

	id, err := records.Insert(p, rec)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}

	resp := make(map[string]string, 1)
	resp["id"] = strconv.FormatUint(id, 10)
	lib.ReturnJSON(w, resp)
}
