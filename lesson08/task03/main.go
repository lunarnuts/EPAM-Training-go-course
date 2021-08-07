package main

import (
	"fmt"
	"log"
	"net/http"

	"com.epam.training/go-course/lesson06/task03/util"
)

func main() {
	port := 8080
	fmt.Printf("Launching on port: %d\n", port)
	http.HandleFunc("/", util.Handler)
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
