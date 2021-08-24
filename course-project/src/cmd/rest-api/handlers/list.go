package user

import (
	"net/http"

	"github.com/lunarnuts/go-course/tree/course-project/cmd/rest-api/lib"
	"github.com/lunarnuts/go-course/tree/course-project/db/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	p := models.ListRequests()
	list, err := p.List()
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	lib.ReturnJSON(w, list)
}
