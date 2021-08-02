package util

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type myError struct {
	Code    int
	Message error
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		switch r.Method {
		case "GET":
			http.ServeFile(w, r, "index.html")
		case "POST":
			name := r.FormValue("name")
			address := r.FormValue("address")
			expiration := time.Now().Add(365 * 24 * time.Hour)
			cookie := http.Cookie{Name: "token", Value: fmt.Sprintf("%s:%s", name, address), Expires: expiration}
			http.SetCookie(w, &cookie)
			http.ServeFile(w, r, "index.html")
		}
	} else {
		errorHandler(w, myError{
			Code:    404,
			Message: errors.New("page not found"),
		})
	}
}

func errorHandler(w http.ResponseWriter, err myError) {
	templ, e := template.ParseFiles("util/error.html")
	if e != nil {
		log.Print(e)
		w.WriteHeader(500)
		templ.Execute(w, myError{
			Code:    500,
			Message: errors.New("Internal server error"),
		})
	}
	w.WriteHeader(err.Code)
	templ.Execute(w, err)
}
