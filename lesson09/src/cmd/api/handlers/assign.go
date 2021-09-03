package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lunarnuts/go-course/tree/lesson09/src/cmd/api/lib"
	"github.com/lunarnuts/go-course/tree/lesson09/src/db"
	"github.com/lunarnuts/go-course/tree/lesson09/src/models"
)

type GroupToContact struct {
	Id  uint64 `json:"id"`
	Gid uint64 `json:"groupId"`
}

func AssignContactToGroup(conn *db.DBConn, w http.ResponseWriter, r *http.Request) {

	var g GroupToContact
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		lib.ReturnClientError(w, err.Error())
		return
	}
	id := g.Id
	gid := g.Gid
	err = models.AssignContactToGroup(*conn, uint64(id), uint64(gid))
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	lib.ReturnJSON(w, fmt.Sprintf("Assigned id:%d to group:%d", id, gid))
}
