package user

import (
	"log"
	"net/http"

	stub_contacts "github.com/wshaman/contacts-stub"

	"github.com/lunarnuts/go-course/lesson07/task01/lib"
)

func SearchFirstName(w http.ResponseWriter, r *http.Request) {
	p := stub_contacts.LoadPersistent()
	list, err := p.List()
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	name := lib.FirstNameFromVars(r)
	log.Print(name)
	result := []stub_contacts.Contact{}
	for _, entity := range list {
		if entity.FirstName == name {
			result = append(result, entity)
		}
	}
	lib.ReturnJSON(w, result)
}
