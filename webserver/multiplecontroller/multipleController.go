package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	rootController()
	helloController()

	log.Print("starting server")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	log.Print("End server")
}

func helloController() {
	http.HandleFunc("/hello", helloHandler)
}

func helloHandler(responseWriter http.ResponseWriter, request *http.Request) {
	log.Print("Inside hello handler")

	for index, header := range request.Header {
		fmt.Fprintf(responseWriter, "Header[%q] = %q\n", index, header)
	}
	fmt.Fprintf(responseWriter, "Host = %q\n", request.Host)
	fmt.Fprintf(responseWriter, "RemoteAddr = %q\n", request.RemoteAddr)
	//writting response

}

func rootController() {
	http.HandleFunc("/", rootHandler)
}

func rootHandler(responseWriter http.ResponseWriter, request *http.Request) {
	log.Print("Inside root handler")
	fmt.Fprintf(responseWriter, "root handler params %q\n %q\n %q\n", request.URL.Path, request.Method, request.Proto)
}
