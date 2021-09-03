package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

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

func ReturnJSON(w http.ResponseWriter, data interface{}) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Unable to encode json: %v", err)
		ReturnInternalError(w, err)
		return
	}
}
