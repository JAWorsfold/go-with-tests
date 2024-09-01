package main

import (
	"log"
	"net/http"

	"go-with-tests/injection"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(injection.MyGreeterHandler)))
}
