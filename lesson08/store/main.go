package main

import (
	"fmt"
	"log"
	"net/http"

	util "github.com/lunarnuts/go-course/tree/lesson08/store/util"
)

func main() {
	port := 8081
	fmt.Printf("Launching on port: %d\n", port)
	http.HandleFunc("/", util.Handler)
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
