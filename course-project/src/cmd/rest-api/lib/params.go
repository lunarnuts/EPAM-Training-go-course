package lib

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func IDFromVars(r *http.Request) (uint64, error) {
	idString := mux.Vars(r)["id"]
	i, err := strconv.Atoi(idString)
	if err != nil {
		return 0, err
	}
	return uint64(i), nil
}

func CityNameFromVars(r *http.Request) string {
	return r.FormValue("city")
}
