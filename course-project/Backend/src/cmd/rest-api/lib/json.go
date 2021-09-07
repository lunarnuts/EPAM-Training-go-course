package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReturnJSON(w http.ResponseWriter, data interface{}) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Unable to encode json: %v", err)
		ReturnInternalError(w, err)
		return
	}
}
