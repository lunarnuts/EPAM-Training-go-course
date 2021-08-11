package util

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"text/template"
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
			formData := url.Values{
				"token": []string{fmt.Sprintf("%s:%s", name, address)}}
			cookie := &http.Cookie{
				Name:  "token",
				Value: name + ":" + address,
			}
			http.SetCookie(w, cookie)
			_, err := http.PostForm("http://localhost:8081", formData)
			if err != nil {
				log.Print(err)
				return
			}
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
		fmt.Fprint(w, "<head><title>500</title></head><body><h1>500</h1><p>Internal Server Error</p></body")
	} else {
		w.WriteHeader(err.Code)
		templ.Execute(w, err)
	}
}
