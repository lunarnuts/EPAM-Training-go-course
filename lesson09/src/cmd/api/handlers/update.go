package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lunarnuts/go-course/tree/lesson09/src/cmd/api/lib"
	"github.com/lunarnuts/go-course/tree/lesson09/src/db"
	"github.com/lunarnuts/go-course/tree/lesson09/src/models"
)

func Update(conn *db.DBConn, w http.ResponseWriter, r *http.Request) {
	var rec models.Contact
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil { // bad request
		lib.ReturnClientError(w, err.Error())
		return
	}

	err = models.UpdateContact(*conn, &rec)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	lib.ReturnJSON(w, fmt.Sprintf("Updated id:%d", rec.ID))
}
