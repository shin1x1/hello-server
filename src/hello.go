package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//noinspection GoUnhandledErrorResult
		fmt.Fprintf(writer, "Hello")
	})
	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		//noinspection GoUnhandledErrorResult
		fmt.Fprintf(writer, "ok")
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
