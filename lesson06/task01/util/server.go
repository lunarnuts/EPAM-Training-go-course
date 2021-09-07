package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type myHeader struct {
	Accept    []string `json:"Accept"`
	UserAgent []string `json:"User_agent"`
}

type myData struct {
	Host       string   `json:"host"`
	UserAgent  string   `json:"user_agent"`
	RequestURI string   `json:"request_uri"`
	Headers    myHeader `json:"headers"`
}

func ResponseHandler(w http.ResponseWriter, r *http.Request) {
	data := myData{
		Host:       r.Host,
		UserAgent:  r.Header.Get("User-Agent"),
		RequestURI: r.RequestURI,
		Headers: myHeader{
			Accept:    strings.Split(r.Header.Get("Accept"), ","),
			UserAgent: strings.Split(r.Header.Get("User-Agent"), "\n"),
		},
	}
	json.NewEncoder(w).Encode(data)
	fmt.Println(r.Header.Get("User-Agent"))
}
