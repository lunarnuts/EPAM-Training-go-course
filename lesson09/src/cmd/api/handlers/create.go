package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lunarnuts/go-course/tree/lesson09/src/cmd/api/lib"
	"github.com/lunarnuts/go-course/tree/lesson09/src/db"
	"github.com/lunarnuts/go-course/tree/lesson09/src/models"
)

func Insert(conn *db.DBConn, w http.ResponseWriter, r *http.Request) {
	var rec models.Contact
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil { // bad request
		lib.ReturnClientError(w, err.Error())
		return
	}

	err = models.InsertContact(*conn, &rec)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	lib.ReturnJSON(w, rec.ID)
}
