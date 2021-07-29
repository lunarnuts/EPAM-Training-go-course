package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type myHeader struct {
	Accept     []string `json:"Accept"`
	User_agent []string `json:"User-Agent"`
}

type myData struct {
	Host        string   `json:"Host"`
	User_agent  string   `json:"UserAgent"`
	Request_uri string   `json:"RequestURI"`
	Headers     myHeader `json:"Header"`
}

func ResponseHandler(w http.ResponseWriter, r *http.Request) {
	data := myData{
		Host:        r.Host,
		User_agent:  r.Header.Get("User-Agent"),
		Request_uri: r.RequestURI,
		Headers: myHeader{
			Accept:     strings.Split(r.Header.Get("Accept"), ","),
			User_agent: strings.Split(r.Header.Get("User-Agent"), "\n"),
		},
	}
	json.NewEncoder(w).Encode(data)
	fmt.Println(r.Header.Get("User-Agent"))
}
