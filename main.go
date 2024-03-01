package main

import (
	"fmt"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "WUBBA LUBBA DUBDUB")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", Home)

	http.ListenAndServe(":8080", nil)
}
