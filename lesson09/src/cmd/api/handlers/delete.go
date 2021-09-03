package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lunarnuts/go-course/tree/lesson09/src/cmd/api/lib"
	"github.com/lunarnuts/go-course/tree/lesson09/src/db"
	"github.com/lunarnuts/go-course/tree/lesson09/src/models"
)

func Delete(conn *db.DBConn, w http.ResponseWriter, r *http.Request) {

	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil { // bad request
		lib.ReturnClientError(w, err.Error())
		return
	}

	err = models.DeleteContact(*conn, uint64(id))
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	lib.ReturnJSON(w, fmt.Sprintf("Deleted id:%d", id))
}
