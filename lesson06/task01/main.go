package main

import (
	"fmt"
	"net/http"

	"com.epam.training/lesson06/util"
)

func main() {
	port := 8080
	fmt.Printf("Listening server on port: %d\n", port)
	http.HandleFunc("/", util.ResponseHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
