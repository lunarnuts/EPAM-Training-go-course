package util

import (
	"encoding/json"
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

type myToken struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json: "createdAt"`
	ExpiredAt time.Time `json: "expiredAt"`
}

var db = make(map[string]interface{})

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		switch r.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(db)
		case "POST":
			token := r.FormValue("token")
			createdAt := time.Now()
			expiredAt := time.Now().Add(10 * 24 * time.Hour)
			temp := myToken{Token: token, CreatedAt: createdAt, ExpiredAt: expiredAt}
			db[token] = temp
			log.Print("token created: ", temp)
			w.WriteHeader(200)
			fmt.Fprint(w, "token stored succesfully:", temp)
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
