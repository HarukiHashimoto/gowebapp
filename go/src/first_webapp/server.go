package main

import (
	f "fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	f.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
