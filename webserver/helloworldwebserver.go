package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	rootController()

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	log.Print("End server")
}

func rootController() {
	http.HandleFunc("/", rootHandler)
	log.Print("starting server")
}

func rootHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "%q\n %q\n %q\n", request.URL.Path, request.Method, request.Proto)
}
