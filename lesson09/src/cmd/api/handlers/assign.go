package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lunarnuts/go-course/tree/lesson09/src/cmd/api/lib"
	"github.com/lunarnuts/go-course/tree/lesson09/src/models"
)

func AssignContactToGroup(conn *sql.DB, w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	gidString := mux.Vars(r)["gid"]
	id, err := strconv.Atoi(idString)
	if err != nil { // bad request
		lib.ReturnClientError(w, err.Error())
		return
	}
	gid, err := strconv.Atoi(gidString)
	if err != nil { // bad request
		lib.ReturnClientError(w, err.Error())
		return
	}

	err = models.AssignContactToGroup(conn, uint64(id), uint64(gid))
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	lib.ReturnJSON(w, fmt.Sprintf("Assigned id:%d to group:%d", id, gid))
}
