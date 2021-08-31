package lib

import (
	"log"
	"net/http"
)

type ErrorEmptyResponse struct{}
type ErrorNotFound struct{}
type ErrorJSONResponse struct{}

func (e *ErrorEmptyResponse) Error() string {
	return "weatherApi - Empty Response"
}

func (e *ErrorNotFound) Error() string {
	return "weatherApi - Not found"
}

func (e *ErrorJSONResponse) Error() string {
	return "weatherApi - JSON error"
}

func ReturnInternalError(w http.ResponseWriter, errHappened error) {
	log.Println(errHappened.Error())
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := w.Write([]byte(`Sorry pal, not this time`)); err != nil {
		log.Fatal(err)
	}
}

func ReturnClientError(w http.ResponseWriter, text string) {
	log.Println(text)
	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write([]byte(text)); err != nil {
		log.Fatal(err)
	}
}
